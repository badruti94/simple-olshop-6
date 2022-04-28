package item

import (
	"net/http"
	"simple-olshop-6/utils/db"

	"github.com/labstack/echo/v4"
)

type Item struct {
	Id          string
	Name        string
	Description string
	Price       int
	Image       string
	Stock       int
	CreatedAt   string `gorm:"column:createdAt"`
	UpdatedAt   string `gorm:"column:updatedAt"`
}

func GetAllItem(c echo.Context) error {
	db, err := db.Connect()

	if err != nil {
		return err
	}

	var items []Item
	db.Find(&items)

	return c.JSON(http.StatusOK, items)
}

func GetItemById(c echo.Context) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	var item Item
	db.First(&item, c.Param("id"))

	return c.JSON(http.StatusOK, item)
}
