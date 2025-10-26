package main

import (
	"Calculator/calc"
	"fmt"
	"strconv"
)

func main() {
	var num1 float64         /// first number
	var num2 float64         /// second number
	var operation string     /// operations + - * /
	var input string         /// a variable created for error handling and filter the user's inputs
	var filterednumb float64 /// variable that sucesfully went trough the filtering process
	var err error
	var finalnumresult float64 ///final result of each calculation, used for every calculations past the first one
	var calculationcheck bool  /// boolean variable that decides in which order should the calculation play out
	var firstOperand float64   /// to make the switch statement cleaner
	var history []string       // a slice of a string, stores all calculations done in the calc
	var index int              // index for the history
	for {
		if calculationcheck == false { /// first for loop, it checks for calculation check, its set to false in default so it executes it first
			fmt.Print("Enter first number: ") /// user input if its exit, exits the entire program,  it scans for a string that isnt exit, continues the loop till a number is the input.
			fmt.Scan(&input)
			if input == "exit" {
				fmt.Println("thank you, come again")
				return
			}
			filterednumb, err = strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Please Enter a valid number")
				continue
			}
			num1 = filterednumb /// after the filtering process num 1 becomes the user's input for future calculations, when the first iteration has been complete. it becomes obsolete

		}
		for {
			fmt.Print("Enter operation (+, -, *, /): ") /// operations, if the user's operation is not equal to any of the listed operations, it'll repeat the message.
			fmt.Scan(&operation)
			if operation == "exit" {
				fmt.Println("thank you, come again") /// the condition of breaking the loop is entering a valid operation, else the loop continues looping.
				return
			}
			if operation == "history" {
				fmt.Println("┌───── Calculation History ───┐")
				fmt.Println(history)
				fmt.Println("└─────────────────────────────┘")
				continue
			}
			if operation == "+" || operation == "/" || operation == "-" || operation == "*" || operation == "exit" {
				break
			} else {
				fmt.Println("Please Enter a valid operation")
				continue
			}
		}

		for {
			fmt.Print("Enter second number: ") ///same as num 1 filters, exit is exiting the program
			fmt.Scan(&input)
			if input == "exit" {
				return
			}
			if input == "history" {
				fmt.Println("┌───── Calculation History ───┐")
				fmt.Println(history)
				fmt.Println("└─────────────────────────────┘")
				continue
			}
			filterednumb, err = strconv.ParseFloat(input, 64) /// and loops when its a string and not a number.
			if err != nil {
				fmt.Println("Please Enter a valid number")
				continue
			}
			break
		}
		num2 = filterednumb ///filtered number 2 for further calculations. unlike num1, its necessary to put it outside the loop for further calculations.

		if calculationcheck == false { /// the boolean logic here is the foundation of everything, its a on and off switch, except in this scenario the switch stays on. u click it on and it stays on, after num 1 and num 2 produce the final result, num 1 becomes entirely obsolete
			firstOperand = num1
			calculationcheck = true ///calculationcheck = true  // Switch to continuous mode - all future calculations use previous result
		} else {
			firstOperand = finalnumresult
		}
		switch operation {
		case "-":
			finalnumresult = calc.Subtract(firstOperand, num2)
			fmt.Println(finalnumresult)
		case "*":
			finalnumresult = calc.Multiply(firstOperand, num2)
			fmt.Println(finalnumresult)
		case "/":
			result, err := calc.Divide(firstOperand, num2)
			if err != nil {
				fmt.Println(err)
				continue
			}
			finalnumresult = result
			fmt.Println(result)
		case "+":
			finalnumresult = calc.Add(firstOperand, num2)
			fmt.Println(finalnumresult)

		default:
			fmt.Println("Unknown operation")
		}
		for index = 0; index < len(history); index++ { // simple condition to increase the number of index proportional to the calculation's done
		}
		calculation := fmt.Sprintf("{%d.  %.2f %s %.2f = %.2f}\n", index, firstOperand, operation, num2, finalnumresult) // sprintf uses the user's input to create the history of calculations.
		history = append(history, calculation)

	} // appends calculation to history, as a string
}

/// this sums it up, finalnumresult takes on multiple values troughout the calculation process, and my calculator.go has all the calculation code.
/// pretty simple code, but it works
