package stringsc

import (
	"fmt"
	"strings"
)

func Strings2() {
	var text = "Hello world, Hello Go"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "Hello", "Hola", -1))
	fmt.Println(strings.Split(text, ","))
}