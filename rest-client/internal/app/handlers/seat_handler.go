package handler

import(
	"github.com/gofiber/fiber/v2"

	model"cinema/pkgs/monorepo"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"
	"api/internal/app/service"
)

func ListSeat(c *fiber.Ctx)error{
	// start service
	data,err := service.ListSeat()
	if err != nil{
		return c.Status(500).JSON(response.Error("list seat",err.Error()))
	}
	// return json data
	return c.Status(200).JSON(response.Pass("list seat", data))
}

func GetSeat(c *fiber.Ctx)error{
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
	data,err := service.GetSeat(id)
	if err != nil{
		return c.Status(500).JSON(response.Error("get seat", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("get seat",data))
}
func CreateSeat(c *fiber.Ctx)error{
	// declared model
	var seat model.CreateSeatParams

	// parser body data to json
	if err := c.BodyParser(&seat);err != nil{
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}

	// validate json
	if err := validate.BodyStructs(seat);err != nil{
		return c.Status(422).JSON(response.Error("validate data", err.Error()))
	}

	// start service
	data,err := service.CreateSeat(seat)
	if err != nil{
		return c.Status(500).JSON(response.Error("create seat",err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("create seat",data))
}
func UpdateSeat(c *fiber.Ctx)error{
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
	var seat model.UpdateSeatParams

	// parser body data to json
	if err := c.BodyParser(&seat);err != nil{
		return c.Status(400).JSON(response.Error("parser json",err.Error()))
	}
	// add id to data
	seat.ID = id

	// validate json
	if err := validate.BodyStructs(seat);err != nil{
		return c.Status(422).JSON(response.Error("validate data", err.Error()))
	}

	// start service
	data,err := service.UpdateSeat(seat)
	if err != nil{
		return c.Status(500).JSON(response.Error("update seat",err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("update seat",data))
}
func DeleteSeat(c *fiber.Ctx)error{
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
	if err := service.DeleteSeat(id);err != nil{
		return c.Status(500).JSON(response.Error("delete seat",err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("delete seat",struct{}{}))
}
