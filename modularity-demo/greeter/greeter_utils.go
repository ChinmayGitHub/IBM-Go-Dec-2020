package greeter

import (
	"fmt"
)

func Greet(name string) string {
	return fmt.Sprintf("Hi %s!", name)
}
