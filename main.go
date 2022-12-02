package main

import (
	"fmt"

	"github.com/alifpramanarasa/go-proto-example/model"
)

func main() {
	var user1 = &model.User{
		Id:       "1",
		Name:     "user1",
		Password: "password1",
		Gender:   model.UserGender_MALE,
	}

	var userList = &model.UserList{
		List: []*model.User{
			user1,
		},
	}

	var garage1 = &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 53.22033123,
		},
	}

	fmt.Println("==== User ====")
	fmt.Printf("Original => %#v", user1)
	fmt.Printf("JSON Content")
	fmt.Println(user1.String())

	fmt.Println("==== User List ====")
	fmt.Printf("Original => %#v", userList)
	fmt.Printf("JSON Content")
	fmt.Println(userList.String())

	fmt.Println("==== Garage ====")
	fmt.Printf("Original => %#v", garage1)
	fmt.Printf("JSON Content")
	fmt.Println(garage1.String())
}
