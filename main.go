// [WIP]
// [TODO]
// * Add support for float operations.
// * Add support for Trignometric operations.
// * Add support for Root operations.
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter text: ")
	var n string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	n = scanner.Text()
	if strings.Contains(n, "+") || strings.Contains(n, "-") || strings.Contains(n, "*") || strings.Contains(n, "/") {
		r := regexp.MustCompile(`[\+\-\*\/]`)
		operator := r.FindAllString(n, -1)[0] // -
		opArr := strings.Split(n, operator)   // ['3 ' ' 10' ]
		op1, _ := strconv.Atoi(strings.Trim(opArr[0], " "))
		op2, _ := strconv.Atoi(strings.Trim(opArr[1], " "))
		switch operator {
		case "+":
			fmt.Println("Adding "+strconv.Itoa(op1)+" + "+strconv.Itoa(op2)+" =", op1+op2)
		case "-":
			fmt.Println("Subtracting "+strconv.Itoa(op1)+" - "+strconv.Itoa(op2)+" =", op1-op2)
		case "*":
			fmt.Println("Multiplying "+strconv.Itoa(op1)+" * "+strconv.Itoa(op2)+" =", op1*op2)
		case "/":
			if op2 == 0 {
				fmt.Println("Cannot divide by 0")
			} else {
				operator1 := float64(op1)
				operator2 := float64(op2)
				fmt.Printf("Dividing %s / %s = %.2f\n", strconv.Itoa(op1), strconv.Itoa(op2), operator1/operator2)
			}

		}
	} else if strings.Contains(n, "sin") || strings.Contains(n, "cos") || strings.Contains(n, "tan") || strings.Contains(n, "cot") || strings.Contains(n, "cosec") || strings.Contains(n, "sec") {
		// Trignometric function
	} else if strings.Contains(n, "sqrt") || strings.Contains(n, "cbrt") {
		// Root function

	}
}
