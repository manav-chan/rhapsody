// what if user could not login or access api
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