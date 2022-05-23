package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductInterface interface {
	create() bool
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

func ServerStart() {
	e := echo.New()
	e.GET("/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Dojo Kun")
	})
	e.POST("/", saveProduct)

	e.Logger.Fatal(e.Start(":8888"))
}

func saveProduct(c echo.Context) error {
	// curl -X POST http://localhost:8888/ -H 'Content-Type: application/json' -d '{"id":"1", "price": 10.20, "description": "Maquininha Sumup"}'

	p := new(Product)
	c.Bind(p)

	fmt.Print(p)

	return c.JSON(http.StatusOK, p)
}

/*
curl --location --request POST 'localhost:8888' \
--header 'Content-Type: application/json' \
--data-raw '{"id": "id-value"}'

*/

/*
POST http://localhost:8888/
Content-Type: multipart/form-data; boundary=WebAppBoundary

--WebAppBoundary
Content-Disposition: form-data; name="id"

id-value
--WebAppBoundary--

*/
