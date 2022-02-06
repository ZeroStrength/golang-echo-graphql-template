package mysql

import (
	"golang-echo-graphql/pkg/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		var u []*model.User

		if err := RW.Find(&u).Error; err != nil {
			// error handling here
			return err
		}

		return c.JSON(http.StatusOK, u)
	}
}
