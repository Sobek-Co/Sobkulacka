package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type ExpTreeNode struct {
	value string
	left  *ExpTreeNode
	right *ExpTreeNode
}

func NewExpTreeNode(token string) *ExpTreeNode {
	return &ExpTreeNode{
		value: token,
	}
}

//************ AI generated code **********************//

// Postorder traversal: Left -> Right -> Root (postfix notation)
func (node *ExpTreeNode) Postorder() string {
	if node == nil {
		return ""
	}

	if node.left == nil && node.right == nil {
		return node.value
	}

	left := node.left.Postorder()
	right := node.right.Postorder()
	result := left
	if left != "" {
		result += " "
	}
	if right != "" {
		result += right + " "
	}
	return result + node.value
}

// PrintTree prints the tree in a visual format
func (node *ExpTreeNode) PrintTree(prefix string, isLeft bool) {
	if node == nil {
		return
	}

	fmt.Print(prefix)
	if isLeft {
		fmt.Print("├── ")
	} else {
		fmt.Print("└── ")
	}
	fmt.Println(node.value)

	if node.left != nil || node.right != nil {
		if node.left != nil {
			newPrefix := prefix
			if isLeft {
				newPrefix += "│   "
			} else {
				newPrefix += "    "
			}
			node.left.PrintTree(newPrefix, true)
		}
		if node.right != nil {
			newPrefix := prefix
			if isLeft {
				newPrefix += "│   "
			} else {
				newPrefix += "    "
			}
			node.right.PrintTree(newPrefix, false)
		}
	}
}

//******************************************************

func create_subtree(node_stack *Stack[*ExpTreeNode], operator string) {
	if operator == "" {
		fmt.Println("Error: Invalid operator")
		return
	}
	op_node := NewExpTreeNode(operator)
	rightNode, errL := node_stack.pop()
	leftNode, errR := node_stack.pop()
	if errL != nil || errR != nil {
		fmt.Println("Error: Invalid expression")
		return
	}
	op_node.left = leftNode
	op_node.right = rightNode
	node_stack.push(op_node)
}

// *** A simple stack implementation using a slice ******
type Stack[T any] []T

func (s *Stack[T]) push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) pop() (T, error) {
	var zero T
	if len(*s) == 0 {
		return zero, fmt.Errorf("stack is empty")
	}
	length := len(*s)
	top := (*s)[length-1]
	*s = (*s)[:length-1]
	return top, nil
}

func (s *Stack[T]) top() (T, error) {
	var zero T
	if len(*s) == 0 {
		return zero, fmt.Errorf("stack is empty")
	}
	return (*s)[len(*s)-1], nil
}

//******************************************************

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

// Tokenize splits the input string into tokens (numbers, operators, parentheses)
// Supports multi-digit and floating-point numbers
func tokenize(expr string) []string {
	// Pattern matches: floating-point numbers, operators, and parentheses
	re := regexp.MustCompile(`\d+\.?\d*|[+\-*/^()]`)
	return re.FindAllString(expr, -1)
}

// IsNumber checks if a token is a number (integer or float)
func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func main1(istring string) {
	// Get input from command-line arguments or use default
	
	
	op_stack := make(Stack[string], 0)
	node_stack := make(Stack[*ExpTreeNode], 0)
	//output := make([]string, 0)
	

	// Tokenize the input
	tokens := tokenize(istring)
	

	parenCount := 0
	// Process each token
	for _, token := range tokens {

		if isNumber(token) {
			// It's a number - add to output
			node_stack.push(NewExpTreeNode(token))
			//output = append(output, token)

		} else if token == "+" || token == "-" || token == "*" || token == "/" || token == "^" {
			// It's an operator
			stackTop, err := op_stack.top()
			for err == nil && top_is_precedent(stackTop, token) {
				topOp, _ := op_stack.pop()
				create_subtree(&node_stack, topOp)
				//output = append(output, topOp)
				stackTop, err = op_stack.top()
			}
			op_stack.push(token)

		} else if token == "(" {
			// Left parenthesis
			op_stack.push(token)
			parenCount++

		} else if token == ")" {
			// Right parenthesis
			if parenCount <= 0 {
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
				create_subtree(&node_stack, topOp)

				//output = append(output, topOp)
			}
			parenCount--
		} else {
			fmt.Printf("Error: Invalid token '%s'\n", token)
			return
		}
	}

	// Pop any remaining operators from the stack to the output
	for {
		topOp, err := op_stack.pop()
		if err != nil {
			break
		}
		if topOp == "(" {
			fmt.Println("Error: Mismatched parentheses")
			return
		}
		create_subtree(&node_stack, topOp)
		//output = append(output, topOp)
	}
	
	//Printing the results
	fmt.Println("Input:", istring)
	fmt.Println("Tokens:", tokens)
	//fmt.Println("Output (postfix):", output) // to show postfix you have to uncomment the lines related to output slice
	

	//*******AI generated code for building the expression tree from postfix tokens*//
	// Get and print the expression tree
	if len(node_stack) > 0 {
		root := node_stack[0]
		fmt.Println("\n=== Expression Tree ===")
		fmt.Println("Visual tree:")
		root.PrintTree("", false)

		fmt.Println("Postfix notation:", root.Postorder())
	}
	//*************************************************************************
}
