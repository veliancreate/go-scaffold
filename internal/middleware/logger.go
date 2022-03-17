package middleware

import "fmt"

type Logger struct {
	Route string
}

func (*Logger) Info(message string) {
	fmt.Println(message)
}

func (*Logger) Error(message string) {
	fmt.Printf("error: %v", message)
}
