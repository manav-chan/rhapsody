package controller

import (
	"fmt" 
	"strconv"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/manav-chan/rhapsody/models"
	"github.com/manav-chan/rhapsody/database"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Unable to parse body")
	}

	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map {
			"message":"Invalid payload",
		})
	}
	return c.JSON(fiber.Map {
		"message":"Post posted successfully",
	})
}

// need to group posts into pages
func AllPost(c *fiber.Ctx) error {
	//We will get pages in the form of a query parameter - ie inside the query only
	page, _ := strconv.Atoi(c.Query("page","1")) // 1 default value
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getblog []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	database.DB.Model(&models.Blog{}).Count(&total)
	return c.JSON(fiber.Map {
		"data":getblog,
		"meta":fiber.Map {
			"total":total, // total pages
			"page":page, // current page
			"last":math.Ceil(float64(int(total)/limit)), // last page
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blogpost)
	return c.JSON(fiber.Map {
		"data": blogpost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog {
		Id:uint(id),
	}

	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body")
	}
	database.DB.Model(&blog).Updates(blog)
	return c.JSON(fiber.Map {
		"message":"Post updated successfully",
	})
}