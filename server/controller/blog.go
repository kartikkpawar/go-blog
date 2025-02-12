package controller

import "github.com/gofiber/fiber/v2"

func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Blog List",
	}
	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Add a blog",
	}
	c.Status(200)
	return c.JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Update Blog",
	}
	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Delete blog for given id",
	}
	c.Status(200)
	return c.JSON(context)
}
