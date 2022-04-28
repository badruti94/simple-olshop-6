package auth

import (
	"net/http"
	"simple-olshop-6/utils/db"
	"simple-olshop-6/utils/time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type M map[string]interface{}

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt string `gorm:"column:createdAt"`
	UpdatedAt string `gorm:"column:updatedAt"`
}

func Register(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	username := c.FormValue("username")
	password := c.FormValue("password")

	passwordByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	password = string(passwordByte)
	time := time.GetTime()
	//err := user.Insert(name, email, username, password)
	db, err := db.Connect()
	if err != nil {
		return err
	}

	db.Select(
		"Name",
		"Email",
		"Username",
		"Password",
		"Role",
		"CreatedAt",
		"UpdatedAt",
	).Create(&User{
		Name:      name,
		Email:     email,
		Username:  username,
		Password:  password,
		Role:      "user",
		CreatedAt: time,
		UpdatedAt: time,
	})

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "registered success")
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	db, err := db.Connect()
	if err != nil {
		return err
	}
	var user User
	db.First(&user, "username = ?", username)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": user.Role,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, M{"user": user, "token": tokenString})
}
