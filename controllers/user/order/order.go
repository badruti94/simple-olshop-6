package order

import (
	"encoding/json"
	"net/http"
	"simple-olshop-6/utils/db"
	"simple-olshop-6/utils/time"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Order struct {
	ID              int
	UserId          int
	Address         string
	PhoneNumber     string
	DeliveryService string
	PaymentMethod   string
	TransferProof   string
	Total           int
	Status          string
	CreatedAt       string `gorm:"column:createdAt"`
	UpdatedAt       string `gorm:"column:updatedAt"`
}

type OrderItem struct {
	ID        int
	OrderId   int
	ItemId    int    `json:"item_id"`
	Qty       int    `json:"qty"`
	CreatedAt string `gorm:"column:createdAt"`
	UpdatedAt string `gorm:"column:updatedAt"`
}

func Orders(c echo.Context) error {
	db, err := db.Connect()

	if err != nil {
		return err
	}

	order := Order{}

	time := time.GetTime()
	order.UserId, _ = strconv.Atoi(c.FormValue("user_id"))
	order.Address = c.FormValue("address")
	order.PhoneNumber = c.FormValue("phone_number")
	order.DeliveryService = c.FormValue("delivery_service")
	order.PaymentMethod = c.FormValue("payment_method")
	order.TransferProof = c.FormValue("transfer_proof")
	order.Total, _ = strconv.Atoi(c.FormValue("total"))
	order.Status = "dikemas"
	order.CreatedAt = time
	order.UpdatedAt = time

	db.Select(
		"UserId",
		"Address",
		"PhoneNumber",
		"DeliveryService",
		"PaymentMethod",
		"TransferProof",
		"Total",
		"CreatedAt",
		"UpdatedAt",
	).Create(&order)

	itemJsonString := c.FormValue("items")
	itemJsonData := []byte(itemJsonString)

	var orderItems []OrderItem
	err = json.Unmarshal(itemJsonData, &orderItems)

	if err != nil {
		return err
	}
	for i := 0; i < len(orderItems); i++ {
		orderItems[i].OrderId = order.ID
		orderItems[i].CreatedAt = time
		orderItems[i].UpdatedAt = time
	}

	db.Select(
		"OrderId",
		"ItemId",
		"Qty",
		"CreatedAt",
		"UpdatedAt",
	).Create(&orderItems)

	return c.JSON(http.StatusOK, "pesanan anda berhasil dipesan")

}

func GetOrder(c echo.Context) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	var order Order
	db.First(&order, "user_id = ?", c.FormValue("user_id"))

	var orderItem []OrderItem
	db.Find(&orderItem, "order_id = ?", order.ID)

	type M map[string]interface{}
	return c.JSON(http.StatusOK, M{"order": order, "order_item": orderItem})

}

func ReceiveOrder(c echo.Context) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	var order Order
	db.First(&order, c.Param("id"))
	order.Status = "diterima"
	db.Save(&order)

	return c.JSON(http.StatusOK, "pesanan diterima")
}
