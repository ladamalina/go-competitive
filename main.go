package main

import "os"

func Solution() {
	// problem solution goes here
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
