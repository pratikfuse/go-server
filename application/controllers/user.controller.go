package controllers

import (
	"fiberapp/database/dao"
	"fiberapp/database/models"

	"fiberapp/application/response"

	"github.com/gofiber/fiber"
)

type UserController struct {
	router *fiber.Group
}

func (user *UserController) InitRoutes() {
	user.router.Post("/login", HandleLogin)
}

func HandleLogin(ctx *fiber.Ctx) {
	var loginUser models.UserLogin
	err := ctx.BodyParser(&loginUser)
	if err != nil {
		ctx.Status(400).JSON(response.ErrorResponse{ErrorMessage: "Invalid payload", Status: 400})
	}
	user, err := dao.Authenticate(loginUser)
	if err != nil {
		ctx.Status(401).JSON(response.ErrorResponse{ErrorMessage: "Invalid username or password", Status: 401})
		return
	}
	ctx.JSON(response.SuccessResponse{Response: user, Status: 200})
}
