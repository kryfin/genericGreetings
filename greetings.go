package greetings

import (
	"fmt"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome YAOOO 2024!", name)
	return message
}
