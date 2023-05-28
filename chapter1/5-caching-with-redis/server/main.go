package main

import (
	"caching-with-redis/routes"
	"caching-with-redis/services"
	"caching-with-redis/utility"
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
