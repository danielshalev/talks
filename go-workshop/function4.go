// decleration
func someFunction() (string, error) {
	return "Hello", nil

}

// call
myString, err := someFunction()

// discard second return value
anotherString, _ := someFunction()