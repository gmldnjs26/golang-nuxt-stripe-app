package controllers

import (
	"context"
	"fmt"
	"net/smtp"

	"ambassador/src/database"
	"ambassador/src/models"

	"github.com/gofiber/fiber/v2"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func Orders(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Preload("OrderItems").Find(&orders)

	for i, order := range orders {
		orders[i].Name = order.FullName()
		orders[i].Total = order.GetTotal()
	}
	return c.JSON(orders)
}

type CreateOrderRequest struct {
	Code      string
	FirstName string
	LastName  string
	Email     string
	Address   string
	Country   string
	City      string
	Zip       string
	Products  []map[string]int
}

func CreateOrder(c *fiber.Ctx) error {
	var request CreateOrderRequest

	if err := c.BodyParser(&request); err != nil {
		fmt.Println(err)
		return err
	}

	link := models.Link{}

	database.DB.Preload("User").First(&link, models.Link{
		Code: request.Code,
	})

	if link.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid link!",
		})
	}

	order := models.Order{
		Code:            link.Code,
		UserId:          link.UserId,
		AmbassadorEmail: link.User.Email,
		FirstName:       request.FirstName,
		LastName:        request.LastName,
		Email:           request.Email,
		Address:         request.Address,
		Country:         request.Country,
		City:            request.City,
		Zip:             request.Zip,
	}

	tx := database.DB.Begin()

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var lineItems []*stripe.CheckoutSessionLineItemParams

	for _, requestProduct := range request.Products {
		product := models.Product{}
		database.DB.First(&product, uint(requestProduct["product_id"]))

		total := product.Price * float64(requestProduct["quantity"])

		item := models.OrderItem{
			OrderId:           order.Id,
			ProductTitle:      product.Title,
			Price:             product.Price,
			Quantity:          uint(requestProduct["quantity"]),
			AmbassadorRevenue: 0.1 * total,
			AdminRevenue:      0.9 * total,
		}

		// database.DB.Create(&item)
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		// For Stripe LineItems
		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			Name:        stripe.String(product.Title),
			Description: stripe.String(product.Description),
			Images:      []*string{stripe.String(product.Image)},
			Amount:      stripe.Int64(100 * int64(product.Price)),
			Currency:    stripe.String("usd"),
			Quantity:    stripe.Int64(int64(requestProduct["quantity"])),
		})
	}

	// Stripe Payment Start
	stripe.Key = "sk_test_51KvyLfCXj5HN8X9OG0WQD9JH5Apwz7LFQjvIYeG8eMCnUS1bA6UUhMZtIiwebnnCHAhIx2rLJtthjJlOwQ3sLKam00FnCljC4r"

	params := stripe.CheckoutSessionParams{
		SuccessURL:         stripe.String("http://localhost:5000/success?source={CHECKOUT_SESSION_ID}"),
		CancelURL:          stripe.String("http://localhost:5000/error"),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems:          lineItems,
	}

	source, err := session.New(&params)
	if err != nil {
		tx.Rollback()
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	order.TransactionId = source.ID
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// Stripe Payment Over

	tx.Commit()

	return c.JSON(order)
}

func CompleteOrder(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	order := models.Order{}

	database.DB.Preload("OrderItems").First(&order, models.Order{
		TransactionId: data["source"],
	})

	if order.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	order.Complete = true
	database.DB.Save(&order)

	// 1.주문을 완료했음으로 랭킹데이터에 링크를 만든 유저(엠버서더)의 price를 높여주는 처리
	// 2.email송신
	// 3.admin ambassador notice this order is completed
	go func(order models.Order) {
		ambassadorRevenue := 0.0
		adminRevenue := 0.0

		for _, item := range order.OrderItems {
			ambassadorRevenue += item.AmbassadorRevenue
			adminRevenue += item.AdminRevenue
		}

		user := models.User{}
		user.Id = order.UserId

		database.DB.First(&user)

		database.Cache.ZIncrBy(context.Background(), "rankings", ambassadorRevenue, user.GetName())

		ambassadorMessage := []byte(fmt.Sprintf("You earned $%f from the link #%s", ambassadorRevenue, order.Code))

		smtp.SendMail("host.docker.internal:1025", nil, "no-reply@email.com", []string{order.AmbassadorEmail}, ambassadorMessage)

		adminMessage := []byte(fmt.Sprintf("Order #%d with a total of $%f has been completed", order.Id, adminRevenue))

		smtp.SendMail("host.docker.internal:1025", nil, "no-reply@email.com", []string{order.AmbassadorEmail}, adminMessage)
	}(order)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// func DeleteOrder(c *fiber.Ctx) error {
// }

// func UpdateOrder(c *fiber.Ctx) error {
// }
