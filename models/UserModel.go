package models

import "fmt"

func getUser() bool  {
	res := InitDatabase()
	fmt.Println("print db connnect res", res)
	return true
}