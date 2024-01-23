package main

import "fmt"

func main() {
	dbConnection := GetDBInstance()
	fmt.Println("Total DB Instance Count: ", dbConnection.instanceCount)

	dbConnection1 := GetDBInstance()
	fmt.Println("Total DB Instance Count: ", dbConnection1.instanceCount)

	dbConnection2 := GetDBInstance()
	fmt.Println("Total DB Instance Count: ", dbConnection2.instanceCount)

}
