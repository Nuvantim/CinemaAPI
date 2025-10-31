package handler

import(
	"github.com/gofiber/fiber/v2"

	model"cinema/pkgs/monorepo"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"
)

func ListCinema(c *fiber.Ctx)error{
	// start service
	data,err := service.ListCinema()
	if err != nil{
		return c.Status(500).JSON(response.Error("list cinema",err.Error()))
	}
	// return json data
	return c.Status(200).JSON(response.Pass("list cinema", data))
}

func GetCinema(c *fiber.Ctx)error{
	// get params id
	params,err := c.Params("id")
	if err != nil{
		return c.Status(400).JSON(response.Error("get id", err.Error()))
	}

	// id validation
	id,err := validate.ValID(id)
	if err != nil{
		return c.Status(400).JSON(response.Error("validation",err.Error()))
	}

	// start service
	data,err := service.GetCinema(id)
	if err != nil{
		return c.Status(500).JSON(response.Error("get cinema", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("get cinema",data))
}
func CreateCinema(c *fiber.Ctx)error{
	// declared model
	var cinema model.CreateCinemaParams

	// parser body data to json
	if err := c.BodyParser(&cinema);err != nil{
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}

	// validate json
	if err := validate.BodyStruct(cinema);err != nil{
		return c.Status(422).JSON(response.Error("validate data", err.Error()))
	}

	// start service
	data,err := service.CreateCinema(cinema)
	if err != nil{
		return c.Status(500).JSON(response.Error("create cinema",err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("create cinema",data))
}
func UpdateCinema(c *fiber.Ctx)error{
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
	var cinema model.UpdateCinemaParams

	// parser body data to json
	if err := c.BodyParser(&cinema);err != nil{
		return c.Status(400).JSON(response.Error("parser json"),err.Error())
	}
	// add id to data
	cinema.ID = id

	// validate json
	if err := validate.BodyStruct(cinema);err != nil{
		return c.Status(422).JSON(response.Error("validate data", err.Error()))
	}

	// start service
	data,err := service.UpdateCinema(cinema)
	if err != nil{
		return c.Status(500).JSON(response.Error("update cinema",err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("update cinema",data))
}
func DeleteCinema(c *fiber.Ctx)error{
	// get params id
	params,err := c.Params("id")
	if err != nil{
		return c.Status(400).JSON(response.Error("get id",err.Error()))
	}

	// id validation
	id,err := validate.ValID(params)
	if err != nil{
		return c.Status(400).JSON(response.Error("validation",err.Error()))
	}

	// start service
	if err := service.DeleteCinema(id);err != nil{
		return c.Status(500).JSON(response.Error("delete cinema",err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("delete cinema",struct{}{}))
}
