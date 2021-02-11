package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	msg := "I wanna pizza"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Закодированное значение: %s\n", encoded)
	fmt.Printf("Раскодированное значение: %s\n", string(decoded))
}
