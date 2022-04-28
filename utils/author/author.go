package author

import (
	"errors"
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Author(c echo.Context, role string) error {
	tokenString := strings.Split(c.Request().Header["Authorization"][0], " ")[1]

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return errors.New("token tidak valid")
	}

	userRole := claims["role"]

	if userRole != role {
		return errors.New("role bukan " + role)
	}

	return nil
}
