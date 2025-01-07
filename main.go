package main

import "os"

var DEBUG bool = false

func Solution() {
	// your code goes here
}

func main() {
	var err error
	bio, err = NewFullBufferedIO(os.Stdin, os.Stdout)
	// bio, err = NewLineBufferedIO(os.Stdin, os.Stdout) // for interactive problems
	if err != nil {
		panic(err.Error())
	}
	defer bio.Flush()

	Solution()
}
