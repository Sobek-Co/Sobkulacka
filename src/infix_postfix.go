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

// createSubtree creates a subtree with an operator and pushes it to the node stack
func createSubtree(nodeStack *Stack[*ExpTreeNode], operator string) error {
	if operator == "" {
		return fmt.Errorf("invalid operator")
	}
	opNode := NewExpTreeNode(operator)
	rightNode, errL := nodeStack.pop()
	leftNode, errR := nodeStack.pop()
	if errL != nil || errR != nil {
		return fmt.Errorf("invalid expression: insufficient operands")
	}
	opNode.left = leftNode
	opNode.right = rightNode
	nodeStack.push(opNode)
	return nil
}

// Stack is a simple stack implementation using a slice
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

// topIsPrecedent checks if the operator on the stack has higher or equal precedence
func topIsPrecedent(topOpStack, currentOp string) bool {
	// Define operator precedence
	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,
		"(": 0,
	}

	p1, exists1 := precedence[topOpStack]
	p2, exists2 := precedence[currentOp]

	if !exists1 || !exists2 {
		return false
	}

	return p1 >= p2
}

// tokenize splits the input string into tokens (numbers, operators, parentheses)
// Supports multi-digit and floating-point numbers
func tokenize(expr string) []string {
	// Pattern matches: floating-point numbers, operators, and parentheses
	re := regexp.MustCompile(`\d+\.?\d*|[+\-*/^()]`)
	return re.FindAllString(expr, -1)
}

// isNumber checks if a token is a number (integer or float)
func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

// parseInfixExpression converts an infix expression to an expression tree
// using the Shunting Yard algorithm. It handles operator precedence and parentheses.
func parseInfixExpression(inputString string) (*ExpTreeNode, error) {
	opStack := make(Stack[string], 0)
	nodeStack := make(Stack[*ExpTreeNode], 0)

	// Tokenize the input
	tokens := tokenize(inputString)

	if len(tokens) == 0 {
		return nil, fmt.Errorf("empty expression")
	}

	parenCount := 0
	// Process each token
	for _, token := range tokens {

		if isNumber(token) {
			// It's a number - add to node stack
			nodeStack.push(NewExpTreeNode(token))

		} else if token == "+" || token == "-" || token == "*" || token == "/" || token == "^" {
			// It's an operator
			stackTop, err := opStack.top()
			for err == nil && topIsPrecedent(stackTop, token) {
				topOp, _ := opStack.pop()
				if err := createSubtree(&nodeStack, topOp); err != nil {
					return nil, err
				}
				stackTop, err = opStack.top()
			}
			opStack.push(token)

		} else if token == "(" {
			// Left parenthesis
			opStack.push(token)
			parenCount++

		} else if token == ")" {
			// Right parenthesis
			if parenCount <= 0 {
				return nil, fmt.Errorf("mismatched parentheses: unexpected closing parenthesis")
			}
			for {
				topOp, err := opStack.pop()
				if err != nil {
					return nil, fmt.Errorf("mismatched parentheses: missing opening parenthesis")
				}
				if topOp == "(" {
					break
				}
				if err := createSubtree(&nodeStack, topOp); err != nil {
					return nil, err
				}
			}
			parenCount--
		} else {
			return nil, fmt.Errorf("invalid token '%s'", token)
		}
	}

	// Pop any remaining operators from the stack
	for {
		topOp, err := opStack.pop()
		if err != nil {
			break
		}
		if topOp == "(" {
			return nil, fmt.Errorf("mismatched parentheses: unclosed opening parenthesis")
		}
		if err := createSubtree(&nodeStack, topOp); err != nil {
			return nil, err
		}
	}

	// Validate that we have exactly one result
	if len(nodeStack) != 1 {
		return nil, fmt.Errorf("invalid expression: malformed input")
	}

	return nodeStack[0], nil
}
