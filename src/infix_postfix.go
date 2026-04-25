package main

import (
	"fmt"
	"strings"
)

// type ExpTreeNode struct {
// 	value string
// 	left  *ExpTreeNode
// 	right *ExpTreeNode
// }

// func NewExpTreeNode(value string) *ExpTreeNode {
// 	return &ExpTreeNode{
// 		value: value,
// 	}
// }

//*** A simple stack implementation using a slice ******
type stack []string

func (s *stack) push(str string) {
	*s = append(*s, str)
}

func (s * stack) pop() (string, error) {
	if len(*s) == 0 {
		return "", fmt.Errorf("stack is empty")
	}
	length := len(*s)
	top := (*s)[length-1]
	*s = (*s)[:length -1]
	return top, nil
}

func (s *stack) top() (string, error) {
	if len(*s) == 0 {
		return "", fmt.Errorf("stack is empty")
	}
	return (*s)[len(*s)-1], nil
}
//******************************************************

//
func top_is_precedent(top_op_stack, current_op string) bool {
	// Define operator precedence
	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,
		"(": 0,
	}

	p1, exists1 := precedence[top_op_stack]
	p2, exists2 := precedence[current_op]

	if !exists1 || !exists2 {
		return false
	}

	return p1 >= p2
}

// const (
// 		Number = iota
// 		Operator
// 		LeftParen 
// 		RightParen
// 	)

// type Token struct {
	
// 	tokenType int
// 	tokenStr string
// }


func main1() {
	istring := "31 + 4 * 2 / ( 1 - 5 )"
	op_stack := make(stack, 0);
	output := make([]string, 0)
	fmt.Println("Input:", istring)
	
	istring = strings.ReplaceAll(istring, " ", "") // Remove spaces for easier processing
	fmt.Println("Input:", istring)

	parenCount := 0
	// Iterate through the input string and process tokens
	for _, token := range istring {

		switch token {

			// Numbers
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				output = append(output, string(token))
				// Can't handle multi-digit numbers

			// Operators
			case '+', '-', '*', '/', '^':
				stackTop, err := op_stack.top()
				for err == nil && top_is_precedent(stackTop, string(token)) {
					topOp, _ := op_stack.pop()
					output = append(output, topOp)
					stackTop, err = op_stack.top()
				}
				op_stack.push(string(token))

			// Parentheses
			case '(':
				op_stack.push(string(token))
				parenCount += 1
			case ')':
				if parenCount<=0 {
					fmt.Println("Error: Mismatched parentheses")
					return
				}
				for {
					topOp, err := op_stack.pop()
					if err != nil {
						fmt.Println("Error: Mismatched parentheses")
						return
					}
					if topOp == "(" {
						break
					}
					output = append(output, topOp)
				}
				parenCount -= 1
			
			default:
				fmt.Printf("Error: Invalid token '%c'\n", token)
				return
			}
	} // end for
	
	//Pop any remaining operators from the stack to the output
	for {
		topOp, err := op_stack.pop()
		if err != nil {
			break
		}
		if topOp == "(" {
			fmt.Println("Error: Mismatched parentheses")
			return
		}
		output = append(output, topOp)
	}
	fmt.Println("Output (postfix):", output)
}