package order

import (
	"net/http"
	"simple-olshop-6/utils/db"

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

func GetOrder(c echo.Context) error {
	db, err := db.Connect()

	if err != nil {
		return err
	}

	var orders []Order
	db.Find(&orders)

	return c.JSON(http.StatusOK, orders)
}

func GetDetailOrder(c echo.Context) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	var order Order
	db.First(&order, c.Param("id"))

	var orderItems []OrderItem
	db.Find(&orderItems, "order_id = ?", order.ID)

	type M map[string]interface{}
	return c.JSON(http.StatusOK, M{"order": order, "order_items": orderItems})
}

func SendOrder(c echo.Context) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	var order Order
	db.First(&order, c.Param("id"))
	order.Status = "dikirim"
	db.Save(&order)

	return c.JSON(http.StatusOK, "pesanan dikirim")
}
