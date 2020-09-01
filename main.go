package main

import "github.com/tarao1006/go_rest_api/app"

func main() {
	a := app.App{}
	a.Initialize("root", "", "go_rest_api_db")

	a.Run(":8080")
}
