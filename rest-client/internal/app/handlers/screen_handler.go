package handler

import(
	"github.com/gofiber/fiber/v2"

	model"cinema/pkgs/monorepo"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"
	"api/internal/app/service"
)

func ListScreen(c *fiber.Ctx)error{
	// start service
	data,err := service.ListScreen()
	if err != nil{
		return c.Status(500).JSON(response.Error("list screen",err.Error()))
	}
	// return json data
	return c.Status(200).JSON(response.Pass("list screen", data))
}

func GetScreen(c *fiber.Ctx)error{
	// get params id
	params,err := c.ParamsInt("id")
	if err != nil{
		return c.Status(400).JSON(response.Error("get id", err.Error()))
	}

	// id validation
	id,err := validate.ValID(params)
	if err != nil{
		return c.Status(400).JSON(response.Error("validation",err.Error()))
	}

	// start service
	data,err := service.GetScreen(id)
	if err != nil{
		return c.Status(500).JSON(response.Error("get screen", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("get screen",data))
}
func CreateScreen(c *fiber.Ctx)error{
	// declared model
	var screen model.CreateScreenParams

	// parser body data to json
	if err := c.BodyParser(&screen);err != nil{
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}

	// validate json
	if err := validate.BodyStructs(screen);err != nil{
		return c.Status(422).JSON(response.Error("validate data", err.Error()))
	}

	// start service
	data,err := service.CreateScreen(screen)
	if err != nil{
		return c.Status(500).JSON(response.Error("create screen",err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("create screen",data))
}
func UpdateScreen(c *fiber.Ctx)error{
	// get params id
	params,err := c.ParamsInt("id")
	if err != nil{
		return c.Status(400).JSON(response.Error("get id",err.Error()))
	}
	
	// id validation
	id,err := validate.ValID(params)
	if err != nil{
		return c.Status(400).JSON(response.Error("valdation",err.Error()))
	}

	// declared model
	var screen model.UpdateScreenParams

	// parser body data to json
	if err := c.BodyParser(&screen);err != nil{
		return c.Status(400).JSON(response.Error("parser json",err.Error()))
	}
	// add id to data
	screen.ID = id

	// validate json
	if err := validate.BodyStructs(screen);err != nil{
		return c.Status(422).JSON(response.Error("validate data", err.Error()))
	}

	// start service
	data,err := service.UpdateScreen(screen)
	if err != nil{
		return c.Status(500).JSON(response.Error("update screen",err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("update screen",data))
}
func DeleteScreen(c *fiber.Ctx)error{
	// get params id
	params,err := c.ParamsInt("id")
	if err != nil{
		return c.Status(400).JSON(response.Error("get id",err.Error()))
	}

	// id validation
	id,err := validate.ValID(params)
	if err != nil{
		return c.Status(400).JSON(response.Error("validation",err.Error()))
	}

	// start service
	if err := service.DeleteScreen(id);err != nil{
		return c.Status(500).JSON(response.Error("delete screen",err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("delete screen",struct{}{}))
}
