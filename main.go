package main

import (
	"github.com/Tatsuemon/go_rest_api/app"
)

func main() {
	a := App{}
	a.Initialize("root", "", "go_rest_api_db")

	a.Run(":8080")
}
