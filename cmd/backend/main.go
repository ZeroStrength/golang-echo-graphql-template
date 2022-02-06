package main

import (
	"golang-echo-graphql/pkg/graphql"
	"golang-echo-graphql/pkg/mysql"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db, _ := mysql.NewDB()
	db.LogMode(true)
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/user", mysql.GetUsers())

	// graphql endpoint here
	h, _ := graphql.NewHandler(db)
	e.POST("/graphql", echo.WrapHandler(h))

	if err := e.Start(":3000"); err != nil {
		log.Fatalln(err)
	}

}
