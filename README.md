# Go Library for Competitive Programming with Generics

Go was once considered challenging to use for competitive programming, but the introduction of generics has significantly improved its suitability and unlocked much more potential for algorithmic problem solving.

This repository provides a collection of Go libraries tailored for competitive programming. It has **no external dependencies** — only the Go standard library is used.

The implementations of algorithms and data structures are inspired by the  
[AtCoder Library](https://github.com/atcoder/ac-library).

## Usage

```go
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
```
