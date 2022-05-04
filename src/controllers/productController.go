package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"ambassador/src/database"
	"ambassador/src/models"

	"github.com/gofiber/fiber/v2"
)

func Products(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)
	log.Printf("%v", product)

	go database.ClearCache("products_frontend", "products_backend")

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var product models.Product

	product.Id = uint(id)

	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{}
	product.Id = uint(id)

	log.Printf("%v", product)

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(&product)

	// go deleteCache("products_frontend") // 비동기처리
	// go func(key string) {               // 이렇게 익명함수로도 go 키워드를 사용할 수 있다.
	// 	time.Sleep(5 * time.Second)
	// 	database.Cache.Del(context.Background(), key)
	// }("products_backend")
	go database.ClearCache("products_frontend", "products_backend")

	return c.JSON(product)
}

// func deleteCache(key string) {
// 	time.Sleep(5 * time.Second) // 이 작업이 5초가 걸린다고 가정
// 	database.Cache.Del(context.Background(), key)
// }

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{}
	product.Id = uint(id)

	database.DB.Delete(&product)

	go database.ClearCache("products_frontend", "products_backend")

	return nil
}

func ProductsFrontend(c *fiber.Ctx) error {
	var products []models.Product
	ctx := context.Background() // 캐시에 필요한 변수

	result, err := database.Cache.Get(ctx, "products_frontend").Result()

	if err != nil { // 캐싱되어 있지 않을때
		fmt.Println(err.Error())

		database.DB.Find(&products)

		bytes, err := json.Marshal(products) // Marshal로 []byte로 변환
		if err != nil {
			panic(err)
		}

		if errKey := database.Cache.Set(ctx, "products_frontend", bytes, 30*time.Minute).Err(); errKey != nil {
			panic(errKey)
		}
	} else {
		json.Unmarshal([]byte(result), &products)
	}

	return c.JSON(products)
}

func ProductsBackend(c *fiber.Ctx) error { // 검색기능
	var products []models.Product
	ctx := context.Background() // 캐시에 필요한 변수

	result, err := database.Cache.Get(ctx, "products_backend").Result()

	if err != nil { // 캐싱되어 있지 않을때
		fmt.Println(err.Error())

		database.DB.Find(&products)

		bytes, err := json.Marshal(products) // Marshal로 []byte로 변환
		if err != nil {
			panic(err)
		}

		if errKey := database.Cache.Set(ctx, "products_backend", bytes, 30*time.Minute).Err(); errKey != nil {
			panic(errKey)
		}
	} else {
		json.Unmarshal([]byte(result), &products)
	}

	var searchedProducts []models.Product

	// search
	if s := c.Query("s"); s != "" {
		lower := strings.ToLower(s)
		for _, product := range products {
			if strings.Contains(strings.ToLower(product.Title), lower) || strings.Contains(strings.ToLower(product.Description), lower) {
				searchedProducts = append(searchedProducts, product)
			}
		}
	} else {
		searchedProducts = products
	}

	// sort
	if sortParam := c.Query("sort"); sortParam != "" {
		sortLower := strings.ToLower(sortParam)
		if sortLower == "asc" {
			sort.Slice(searchedProducts, func(i, j int) bool {
				return searchedProducts[i].Price < searchedProducts[j].Price
			})
		} else if sortLower == "desc" {
			sort.Slice(searchedProducts, func(i, j int) bool {
				return searchedProducts[i].Price > searchedProducts[j].Price
			})
		}
	}

	// pagination
	total := len(searchedProducts)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage := 9

	var data []models.Product

	if total <= page*perPage && total >= (page-1)*perPage {
		data = searchedProducts[(page-1)*perPage : total]
	} else if total >= page*perPage {
		data = searchedProducts[(page-1)*perPage : page*perPage]
	} else {
		data = []models.Product{}
	}

	return c.JSON(fiber.Map{
		"data":      data[(page-1)*perPage : page*perPage],
		"total":     total,
		"page":      page,
		"last_page": total/perPage + 1,
	})
}
