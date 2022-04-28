package main

import (
	"log"
	"os"
	adminItem "simple-olshop-6/controllers/admin/item"
	adminOrder "simple-olshop-6/controllers/admin/order"
	"simple-olshop-6/controllers/auth"
	"simple-olshop-6/controllers/item"
	"simple-olshop-6/controllers/user/order"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	r := echo.New()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r.POST("/auth/register", auth.Register)
	r.POST("/auth/login", auth.Login)

	r.GET("/item", item.GetAllItem)
	r.GET("/item/:id", item.GetItemById)

	adminRoute := r.Group("admin", func(hf echo.HandlerFunc) echo.HandlerFunc {
		return hf
	})
	adminRoute.GET("/item", adminItem.GetAllItem)
	adminRoute.GET("/item/:id", adminItem.GetItemById)
	adminRoute.POST("/item", adminItem.Insert)
	adminRoute.PUT("/item/:id", adminItem.Update)
	adminRoute.DELETE("/item/:id", adminItem.Delete)

	adminRoute.GET("/order", adminOrder.GetOrder)
	adminRoute.GET("/order/:id", adminOrder.GetDetailOrder)
	adminRoute.GET("/order/:id/send", adminOrder.SendOrder)

	userRoute := r.Group("user", func(hf echo.HandlerFunc) echo.HandlerFunc {
		return hf
	})
	userRoute.POST("/order", order.Orders)
	userRoute.GET("/order", order.GetOrder)
	userRoute.GET("/order/:id/receive", order.ReceiveOrder)

	port := os.Getenv("PORT")

	r.Start(":" + port)

}
