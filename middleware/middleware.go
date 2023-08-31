// acts as a gatekeeper that ensures users are authenticated and authorized before accessing certain parts of a web application or API.
package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manav-chan/rhapsody/util"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt") // jwt token head
	if _, err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map {
			"message":"Unauthenticated",
		})
	}
	return c.Next() // next route to run, next function
}