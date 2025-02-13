package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kartikkpawar/go-blog/database"
	"github.com/kartikkpawar/go-blog/model"
)

func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Blog List",
	}
	c.Status(200)

	db := database.DBConnection

	var records []model.Blog

	db.Find(&records)
	context["blog_records"] = records

	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Add a blog",
	}

	record := new(model.Blog)

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request")
	}

	result := database.DBConnection.Create(record)

	if result.Error != nil {
		log.Println("Error saving data")
		context["statusText"] = ""
		context["msg"] = "Something went wrong"
		return c.JSON(context)
	}

	context["data"] = record
	context["msg"] = "Data saved successfully"

	c.Status(201)
	return c.JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Update Blog",
	}
	c.Status(200)

	blogId := c.Params("blogId")

	var record model.Blog
	database.DBConnection.First(&record, blogId)

	if record.ID == 0 {
		log.Println("Record Not Found")
		context["message"] = "Record not found"
		context["statusText"] = ""
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request")
	}

	result := database.DBConnection.Save(record)

	if result.Error != nil {
		log.Println("Error updating data")
		context["message"] = "Something went wrong"
		context["statusText"] = ""
		return c.JSON(context)
	}

	context["data"] = record
	context["message"] = "Record updated successfully"

	c.Status(200)

	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Delete blog for given id",
	}

	blogId := c.Params("blogId")

	var record model.Blog

	database.DBConnection.First(&record, blogId)

	if record.ID == 0 {
		context["message"] = "Record not found"
		context["statusText"] = ""
		c.Status(400)
		return c.JSON(context)
	}

	result := database.DBConnection.Delete(&record)

	if result.Error != nil {
		log.Fatal("Unable to delete")
		context["message"] = "Unable to delete"
		context["statusText"] = ""
		c.Status(400)
		return c.JSON(context)
	}

	context["message"] = "Record Delete Successfully"

	c.Status(200)
	return c.JSON(context)
}
