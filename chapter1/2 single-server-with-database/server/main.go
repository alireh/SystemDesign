package main

import (
	"chapter1-single-server-with-database/routes"
	"chapter1-single-server-with-database/services"
	"chapter1-single-server-with-database/utility"
	"log"
	"net/http"
)

func main() {
	var db = utility.GetConnection()
	services.SetDB(db)
	var appRouter = routes.CreateRouter()

	log.Println("Listening on Port 8002")
	log.Fatal(http.ListenAndServe(":8002", appRouter))
}

//https://medium.com/@yashaswi_nayak/go-with-orm-into-the-g-orm-hole-450025ea0fd8
