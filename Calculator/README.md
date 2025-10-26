# Go Calculator

A command-line calculator built in Go with continuous calculation mode and history tracking.

## Features

- **Basic Operations**: Addition, subtraction, multiplication, division
- **Continuous Calculation Mode**: Results automatically become the first operand for the next calculation
- **Calculation History**: View all past calculations with the `history` command
- **Input Validation**: Handles invalid inputs and division by zero
- **Exit Anytime**: Type `exit` at any prompt to quit
- **Chained Input**: Enter multiple operations in sequence (e.g., 25 + 50 + 22 / 2) (altough it might look unsightly)

## Installation
```bash
git clone https://github.com/Savoxero/GolangLearningLog/Projects/Calculator.git
next step:
locate the Calculator directory and run:
go run main.go/go run all
```

## Usage
```
Enter first number: 10
Enter operation (+, -, *, /): +
Enter second number: 5
15

Enter operation (+, -, *, /): *
Enter second number: 2
30

Enter operation (+, -, *, /): history
┌───── Calculation History ───┐
{1.  10.00 + 5.00 = 15.00}
{2.  15.00 * 2.00 = 30.00}
└─────────────────────────────┘
```

## Special Commands

- `history` - View all calculations performed in current session
- `exit` - Exit the calculator

## Project Structure
```
Calculator/
├── main.go           # Main calculator logic and user interface
└── calc/
    └── operations.go # Mathematical operations (add, subtract, multiply, divide)
```

## Error Handling

- Invalid number inputs prompt for re-entry
- Division by zero displays error and continues
- Invalid operations prompt for valid operator

## Future Improvements

- [ ] Add scientific functions (sqrt, power, sin, cos)
- [ ] Expression parsing for single-line input
- [ ] Save/load history to file
- [ ] Variable storage (x = 10, then use x)

## Learning Goals

This project demonstrates:
- Function organization and packages
- User input handling and validation
- Error handling patterns
- State management with boolean flags
- Slice operations for data storage
- String formatting for display

---

Built while learning Go through boot.dev