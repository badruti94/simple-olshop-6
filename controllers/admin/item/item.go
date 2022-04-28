package item

import (
	"fmt"
	"net/http"
	"simple-olshop-6/utils/db"
	"simple-olshop-6/utils/time"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Item struct {
	ID          int
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

func Insert(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	image := c.FormValue("image")
	price := c.FormValue("price")
	stock := c.FormValue("stock")

	db, err := db.Connect()
	if err != nil {
		return err
	}

	priceString, err := strconv.Atoi(price)
	if err != nil {
		return nil
	}
	stockString, err := strconv.Atoi(stock)
	if err != nil {
		return nil
	}
	time := time.GetTime()

	item := Item{
		Name:        name,
		Description: description,
		Image:       image,
		Price:       priceString,
		Stock:       stockString,
		CreatedAt:   time,
		UpdatedAt:   time,
	}

	db.Select(
		"Name",
		"Description",
		"Price",
		"Image",
		"Stock",
		"CreatedAt",
		"UpdatedAt",
	).Create(&item)

	fmt.Println(item)

	return c.JSON(http.StatusOK, "item berhasil disimpan")
}

func Update(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	image := c.FormValue("image")
	price := c.FormValue("price")
	stock := c.FormValue("stock")

	db, err := db.Connect()

	if err != nil {
		return err
	}

	var item Item
	db.First(&item, c.Param("id"))

	item.Name = name
	item.Description = description
	item.Image = image
	item.Price, _ = strconv.Atoi(price)
	item.Stock, _ = strconv.Atoi(stock)

	db.Save(&item)

	return c.JSON(http.StatusOK, "item berhasil diupdate")
}

func Delete(c echo.Context) error {
	db, err := db.Connect()
	if err != nil {
		return nil
	}

	var item Item
	item.ID, _ = strconv.Atoi(c.Param("id"))

	db.Delete(&item)

	return c.JSON(http.StatusOK, "item berhasil dihapus")
}
