package controller

import (
	"fmt"
	"log"
	"strings"
	"regexp"
	"time"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/manav-chan/rhapsody/models"
	"github.com/manav-chan/rhapsody/database"
	"github.com/manav-chan/rhapsody/util"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)

}

// *fiber.Ctx is a context object which carries information about the HTTP request as well as methods to read and write response data, metadata.
func Register(c *fiber.Ctx) error {

	//When you need to store a collection of arbitrary values of any type, then, identified by strings, a map[string]interface{} or map[string]any is the ideal choice.
	var data map[string] interface{}
	var userData models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body");
	}

	// check length of password - less than 6 invalid
	// convert to string as password is array of bytes
	if len(data["password"].(string)) < 6 { 
		c.Status(400)
		return c.JSON(fiber.Map {
			"message" : "Password must be atleast 6 character long",
		})
	}

	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map {
			"message" : "Invalid email address",
		})
	}

	//userData is a destination struct
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map {
			"message" : "Email already exists",
		})
	}

	user := models.User {
		FirstName: data["first_name"].(string),
		LastName: data["last_name"].(string),
		Email: strings.TrimSpace(data["email"].(string)),
	}

	//encrypting password using bycrypt
	user.SetPassword(data["password"].(string))

	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}

	c.Status(200)
	return c.JSON(fiber.Map {
		"user" : user,
		"message" : "Account Created Successfully",
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string] string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	
	// checks if user's email is present in db or not
	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message" : "Email address does not exist, kindly createa an account",
		})
	}

	// checks if password is correct or not
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400) // Bad Request
		return c.JSON(fiber.Map {
			"message" : "Incorrect password",
		})
	}

	// if correct - authenticate user
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)), )
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	cookie := fiber.Cookie {
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24), // expires after 24 hours
		HTTPOnly: true,
	}
	
	c.Cookie(&cookie)
	return c.JSON(fiber.Map {
		"message":"you have successfully logged in",
		"user":user,
	})
}

type Claims struct {
	jwt.StandardClaims
}