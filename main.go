package main

func main() {
	a := App{}
	a.Initialize("root", "", "go_rest_api_db")

	a.Run(":8080")
}
