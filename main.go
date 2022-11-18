package main

import (
	"dans/routes"
	"dans/utils"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	var port string

	flag.StringVar(&port, "port", os.Getenv("PORT"), "port of the service")

	db := utils.GetDBConnection()
	defer db.Close()

	routes := &routes.Routes{DB:db}

	routes.Setup(port)

}

