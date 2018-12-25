package main

import "fmt"

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)

	if l <= 0 {
		return s, ""
	}

	return s[:l-1], s[l-1]
}

func evalBalance(cur string, prev string) bool {

	if prev == "" {
		return false
	}

	switch cur {
	case "]":
		return prev == "["
	case ")":
		return prev == "("
	case "}":
		return prev == "{"
	default:
		return false
	}
}

func isOpenBracket(cur string) bool {
	return cur == "[" || cur == "{" || cur == "("
}

func isClosedBracket(cur string) bool {
	return cur == "]" || cur == "}" || cur == ")"
}

// Complete the isBalanced function below.
func isBalanced(s string) string {

	if len(s) < 1 {
		return "NO"
	}

	var stac stack
	var last string
	stac = stack{}

	for _, k := range s {
		// fmt.Println("Stack now ", stac)
		cur := string(k)
		// fmt.Println("char", cur)

		if isOpenBracket(cur) {
			stac = stac.Push(cur)
		}

		if isClosedBracket(cur) {

			if len(stac) < 1 {
				return "NO"
			}

			stac, last = stac.Pop()
			if !evalBalance(cur, last) {
				return "NO"
			}
		}

	}

	if len(stac) > 0 {
		return "NO"
	}
	return "YES"
}

func main() {
	var input, output string
	input = "{[()]}"
	output = isBalanced(input)
	fmt.Println(output)
	input = "{[(])}"
	output = isBalanced(input)
	fmt.Println(output)
	input = "{{[[(())]]}}"
	output = isBalanced(input)
	fmt.Println(output)

	input = ")), ]], }}"
	output = isBalanced(input)
	fmt.Println(output)
}
