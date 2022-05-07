package middlewares

import (
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

const SecretKey = "secret"

type ClaimsWithScope struct {
	jwt.StandardClaims // 이 안에 있는 요소들이 distructuring 하는것 처럼 풀어서 들어가는것 같다.
	Scope              string
}

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	payload := token.Claims.(*ClaimsWithScope)
	IsAmbassador := strings.Contains(c.Path(), "/api/ambassador")

	if (payload.Scope == "admin" && IsAmbassador) || (payload.Scope == "ambassador" && !IsAmbassador) {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	return c.Next()
}

func GetUserId(c *fiber.Ctx) (uint, error) {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return 0, err
	}

	payload := token.Claims.(*ClaimsWithScope)

	id, _ := strconv.Atoi(payload.Subject)

	return uint(id), nil
}

func GenerateJWT(userId uint, scope string) (string, error) {
	payload := ClaimsWithScope{}
	payload.Subject = strconv.Itoa(int(userId)) // 이거랑 string(user.Id)랑 다른가? 보다
	payload.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	payload.Scope = scope

	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(SecretKey))
}
