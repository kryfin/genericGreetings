package greetings

import (
	"fmt"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome NEw mEmber!", name)
	return message
}
