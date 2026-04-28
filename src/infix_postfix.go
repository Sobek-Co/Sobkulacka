package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type ExpTreeNode struct {
	value   string
	left    *ExpTreeNode
	right   *ExpTreeNode
	isUnary bool
	unaryOp string
}

func NewExpTreeNode(token string) *ExpTreeNode {
	return &ExpTreeNode{
		value: token,
	}
}

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

	// Check if it's a unary operator
	if isUnaryPrefix(operator) {
		// Unary prefix operator (sin, cos, tan, cot) - pop one operand
		operand, err := nodeStack.pop()
		if err != nil {
			return fmt.Errorf("invalid expression: %s requires an operand", operator)
		}
		opNode.left = operand
		opNode.isUnary = true
		opNode.unaryOp = operator
	} else if operator == "!" {
		// Postfix unary operator (factorial) - pop one operand
		operand, err := nodeStack.pop()
		if err != nil {
			return fmt.Errorf("invalid expression: factorial requires an operand")
		}
		opNode.left = operand
		opNode.isUnary = true
		opNode.unaryOp = "!"
	} else {
		// Binary operator - pop two operands
		rightNode, errL := nodeStack.pop()
		leftNode, errR := nodeStack.pop()
		if errL != nil || errR != nil {
			return fmt.Errorf("invalid expression: insufficient operands")
		}
		opNode.left = leftNode
		opNode.right = rightNode
	}

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

// isUnaryPrefix checks if a token is a prefix unary operator
func isUnaryPrefix(token string) bool {
	switch token {
	case "sin", "cos", "tan", "cot":
		return true
	default:
		return false
	}
}

// operatorPrecedence returns the precedence of an operator
// Higher number means higher precedence
func operatorPrecedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/", "%":
		return 2
	case "^":
		return 3
	case "sin", "cos", "tan", "cot":
		return 4 // Unary functions have high precedence
	case "!":
		return 5 // Factorial has highest precedence
	default:
		return 0
	}
}

// topIsPrecedent checks if the operator on the stack has higher or equal precedence
func topIsPrecedent(topOpStack, currentOp string) bool {
	p1 := operatorPrecedence(topOpStack)
	p2 := operatorPrecedence(currentOp)

	if p1 == 0 || p2 == 0 {
		return false
	}

	return p1 >= p2
}

// tokenize splits the input string into tokens
// Pattern matches: floating-point numbers, function names, operators, and parentheses
func tokenize(expr string) []string {
	re := regexp.MustCompile(`\d+\.?\d*|sin|cos|tan|cot|[+\-*/%^()!]`)
	return re.FindAllString(expr, -1)
}

// isNumber checks if a token is a number (integer or float)
func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

// parseInfixExpression converts an infix expression to an expression tree
// using the Shunting Yard algorithm with support for unary and postfix operators
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

		} else if isUnaryPrefix(token) {
			// Prefix unary operator (sin, cos, tan, cot)
			// Push to operator stack with high precedence
			opStack.push(token)

		} else if token == "!" {
			// Postfix unary operator (factorial)
			// Apply to the top node in the stack immediately
			operand, err := nodeStack.pop()
			if err != nil {
				return nil, fmt.Errorf("invalid expression: factorial operator needs an operand")
			}
			factNode := NewExpTreeNode("!")
			factNode.left = operand
			factNode.isUnary = true
			factNode.unaryOp = "!"
			nodeStack.push(factNode)

		} else if token == "+" || token == "-" || token == "*" || token == "/" || token == "^" || token == "%" {
			// Binary operator
			stackTop, err := opStack.top()
			for err == nil && stackTop != "(" && topIsPrecedent(stackTop, token) {
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

			// Pop operators until we see "("
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

			// Check if there's a function on top of the stack
			peek, err := opStack.top()
			if err == nil && isUnaryPrefix(peek) {
				fn, _ := opStack.pop()
				if err := createSubtree(&nodeStack, fn); err != nil {
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

		// Handle unary prefix operators that weren't used - this is an error
		if isUnaryPrefix(topOp) {
			return nil, fmt.Errorf("incomplete function call: %s requires an argument", topOp)
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
