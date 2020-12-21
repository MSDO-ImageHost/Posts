package main

import (
	"fmt"

	"github.com/MSDO-ImageHost/Posts/internal/auth"
	"github.com/MSDO-ImageHost/Posts/internal/utils"
)

func main() {

	tokenString := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIwIiwicm9sZSI6MCwiaXNzIjoiSW1hZ2VIb3N0LnNkdS5kayIsImlhdCI6MTYwODUxNjYzMiwiZXhwIjoxNjExMzYwMDAwfQ.-yI4mlTI7it8SS0youU_VG8b3Bju7wAz573SUQR9ntA"
	userAuth, err := auth.AuthJWT(tokenString)
	fmt.Println(utils.PrettyFormatMap(userAuth), err)

	tokenString = "eyJ0eXAiOiJKV1QiLCJhbdciOiJIUzI1NiJ9.eyJzdWIiOiIwIiwicm9sZSI6MCwiaXNzIjoiSW1hZ2VIb3N0LnNkdS5kayIsImlhdCI6MTYwODUxNjYzMiwiZXhwIjoxNjExMzYwMDAwfQ.-yI4mlTI7it8SS0youU_VG8b3Bju7wAz573SUQR9ntA"
	userAuth, err = auth.AuthJWT(tokenString)
	fmt.Println(utils.PrettyFormatMap(userAuth), err)

	fmt.Println([]byte("Hello World!"))
}
