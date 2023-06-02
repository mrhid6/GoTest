package main

import (
	"fmt"
	"github.com/mrhid6/GoTest/db"
)

func main(){
	fmt.Println("Hello World")
	fmt.Println(db.testMongo());
}