package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Status int

const (
	NotStarted Status = iota + 1
	Reading
	Finished
)

type Book struct {
	Name        string
	Author      string
	Description string
	Status      Status
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is running")
	})

	e.GET("/books", getAllBooks)

	e.Logger.Fatal(e.Start(":1323"))
}

func getAllBooks(c echo.Context) error {

	books := []Book{
		{
			Name:        "And Then There Were None",
			Author:      "Agatha Christie",
			Description: "People just died at an misterious island.",
			Status:      Finished,
		},
		{
			Name:   "Flowers for Algernon",
			Author: "Daniel Keyes",
			Status: NotStarted,
		},
	}

	return c.JSON(http.StatusOK, struct{ Books []Book }{Books: books})
}
