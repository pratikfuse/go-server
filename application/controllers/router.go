package controllers

import (
	"github.com/gofiber/fiber"
)

type Router struct {
	App *fiber.App
}

func (r *Router) InitRoutes() {
	userRouter := &UserController{router: r.App.Group("/auth")}
	userRouter.InitRoutes()
}
