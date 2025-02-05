//stack implementation in Go generated by ChatGPT
package main

import ("fmt"
)


// Stack structure of size 100 to hold the stack data
type Stack struct {
	elements [100]int
	top int
}

// Push adds an element to the stack
func (s *Stack) Push(element int) bool {
   if s.top == 100 { // if the stack if full the push will fail
	return false
}

s.elements[s.top] = element // set the top element to the value provided
s.top = s.top + 1 // increase top element by 1
return true //return true since the push was successful
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() (int, bool) {
    if s.top == 0 {
        return 0, false // Return 0 and false if the stack is empty
    }
    s.top = s.top - 1 // decrement top element by 1
    return s.elements[s.top], true // return the new top element of the stack and true since the pop was successful
}

// Top returns the top element without removing it from the stack
func (s *Stack) Top() (int, bool) {
    if s.top == 0 { // return 0 and false because the stack is empty
        return 0, false
    }
    return s.elements[s.top -1], true // return the top element of the stack and true since the function executed successfully
}

func main() {
	// Create a new stack
	stack := Stack{}

	// Push some elements onto the stack
	stack.Push(1)
	fmt.Println("Pushed 1")
	stack.Push(2)
	fmt.Println("Pushed 2")
	stack.Push(3)
	fmt.Println("Pushed 3")

	// set top1 to the top element of the stack
	top1, _ := stack.Top()
	fmt.Println("Top element:", top1)
	//pop the top element from the stack
	popped1, _ := stack.Pop()
	fmt.Println("Popped: ", popped1)
	//push some more elements to the stack
	stack.Push(17)
	fmt.Println("Pushed 17")
	stack.Push(19)
	fmt.Println("Pushed 19")
	//pop the top elements from the stack and print them
	popped2, _ := stack.Pop()
	fmt.Println("Popped: ", popped2)
	popped3, _ := stack.Pop()
	fmt.Println("Popped: ", popped3)
	popped4, _ := stack.Pop()
	fmt.Println("Popped: ", popped4)
	//print the top element of the stack
	top2, _ := stack.Top()
	fmt.Println("Top element:", top2)
}
