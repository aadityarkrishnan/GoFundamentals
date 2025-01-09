package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name                 string
	Phone                string
	ScheduledForDeletion bool
}

func removeUser(userMap map[string]User, name string) (map[string]User, error) {
	existingUser, ok := userMap[name]
	if !ok {
		return userMap, errors.New("user does not exist")
	}
	if existingUser.ScheduledForDeletion {
		delete(userMap, name)
	}
	return userMap, nil
}

func main() {
	u1 := User{
		Name:                 "John Doe",
		Phone:                "111111111",
		ScheduledForDeletion: true,
	}

	u2 := User{
		Name:                 "Jane Doe",
		Phone:                "1111111111",
		ScheduledForDeletion: false,
	}

	// Initialize the map using a literal
	userMap := map[string]User{
		"John Doe": u1,
		"Jane Doe": u2,
	}

	userMap, err := removeUser(userMap, "John Doe")
	if err != nil {
		panic(err)
	}

	for key, value := range userMap {
		fmt.Println(key, value)
	}
}
