package main

import (
	"fmt"

	"github.com/chhandakguria/GoAPIs/Packages/Calc"
	"github.com/chhandakguria/GoAPIs/Packages/Numbers"
)

func main() {

	fmt.Println("Result: ", Calc.Calculator(10, 5, '/'))
	fmt.Println("Staus: ", Numbers.IsEven(40))

}
