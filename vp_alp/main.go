package main

import(
	"vp_alp/db"
	"vp_alp/routes"
	
)


func main() {
	db.Init()
	e := routes.Init()
	
	e.Logger.Fatal(e.Start(":7070"))
}
