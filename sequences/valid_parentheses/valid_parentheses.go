package main

import (
	"fmt"
	"strings"
)
func main(){
	text := "{{}}()[()]"
	result := isValid(text)
	fmt.Println(result)
}

func isValid(s string) bool {
	a := strings.Split(s, "")
	myStack := Stack{}
	for i:=0; i< len(s); i++{
		if a[i] == "{" || a[i] == "(" || a[i] == "[" {
			myStack.Push(a[i])
		}else{	
			//if we encounter a closing bracket and no pre-existing open bracket, 
			//then it is invalid
			if len(myStack.items) == 0 {
				return false
			}else{
				top_of_stack := myStack.items[len(myStack.items)-1]
				if (a[i] == "}" && top_of_stack == "{") || (a[i] == ")" && top_of_stack == "(") || (a[i] == "]"&& top_of_stack == "["){
					myStack.Pop()
				}else{
					return false
				}
			}
		}
	}
	if len(myStack.items) == 0{
		return true
	}else if !(len(myStack.items) == 0){
		return false
	}
return false
}

// stack represents stack that holds a slice
type Stack struct {
	items []string
}

//push adds a value to the stack
func (s *Stack) Push(i string) {
	s.items = append(s.items, i)
}

//pop removes a value from the stack and returns value removed
func (s *Stack) Pop() string{
	last := len(s.items)-1
	toRemove := s.items[last]
	s.items = s.items[:last]
	return toRemove
}