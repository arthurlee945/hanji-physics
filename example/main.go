package main

import (
	"fmt"

	"github.com/arthurlee945/suhag/vec"
)

func main() {
	//TESTING GROUND WILL BE REMOVED
	fmt.Println(vec.Sub(*vec.NewVec3(3, 10, 25), *vec.NewVec3(142, 32, 25)))
	fmt.Println(vec.Div(*vec.NewVec2(-3.658830483394366e-05, -9.306608377587624e-05), 10))
}
