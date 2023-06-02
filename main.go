package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/mrhid6/GoTest/routes"

	"github.com/rahul-sinha1908/go-mongoose/mongoose"

	"github.com/gin-gonic/gin"
)

func main(){

	fmt.Println("Starting App!")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	mongoPort, err := strconv.Atoi(os.Getenv("MONGODB_PORT"))

	if(err != nil){
		panic(err);
	}

	mongoose.InitiateDB(mongoose.DBConnection{
		Host: os.Getenv("MONGODB_HOST"),
		Port: mongoPort,
		Database: os.Getenv("MONGODB_DATABASE"),
	});


	router := gin.Default()

	apiv1:=router.Group("/api/v1")
	routes.AddAccountRoutes(apiv1);
	
    router.Run(os.Getenv("HTTP_BIND"));
}