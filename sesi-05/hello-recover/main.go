package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("======================================================")
	// Recover

	defer catchErr()

	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

// func menangkap error
func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("APllication running perfectly")
	}
}

// func validation password
func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("Password has to have more than 4 characters")
	}

	return "Valid password", nil
}
