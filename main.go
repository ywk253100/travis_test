package main

import (
	"fmt"
	"log"
	"travis_test/dao"
)

func main() {
	users, err := dao.GetAllUser()
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Println(users)
}
