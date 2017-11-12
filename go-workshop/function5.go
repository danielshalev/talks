// decleration
func someFunction() (string, error) {
	return "Hello", nil

}

// call
myString, err := someFunction()

if err != nil {
	// Handle error
}