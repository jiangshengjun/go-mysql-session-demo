package main

import (
	"fmt"
	"net/http"

	"github.com/jiangshengjun/go-mysql-session"
	_ "github.com/jiangshengjun/go-mysql-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middlewaresss
	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	defer db.Close()

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		data, err := db.Session("zhuanqian_1").Select("select * from lottery limit 10")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(len(data))

		// for _, row := range data {
		// 	id, _ := strconv.Atoi(row["id"])
		// 	// fmt.Print(id)
		// 	// fmt.Print("|")
		// 	// fmt.Println(row["created_at"])
		// }

		return c.String(http.StatusOK, "Hello, World!\n")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
