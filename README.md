# What is vptr ? 
This package provides a set of functions for converting values to pointers.

There are scenarios when you have to use a pointer to a value without creating a variable. For example


## Installation:

```
go get github.com/alidevhere/vptr
```

# Problem
This package does not convert types, it just converts values to pointers.
It is useful for converting values to pointers in the function parameters, without making code ugly.
For example:
This is how normally one would pass pointer params to a function.
```
func f(a *int, b *string, c *bool) {
	...
}

func main() {
	a := 1
	b := "2"
	c := true
	f(&a, &b, &c)
}
```

Or if someone might do it this way:

```
func f(a *int, b *string, c *bool) {
	...
}

func main() {

	f(&[]int{1}[0], &[]string{"2"}[0], &[]bool{true}[0])
}

```
# Solution:

But anyway it makes code ugly and verbose very quickly.
So there is an easy way to do it too.


It is possible to use this package to convert values to pointers:
```

"github.com/alidevhere/vptr"

func f(a *int, b *string, c *bool) {
	...
}

func main() {
	f(vptr.Int(1), vptr.String("2"), vptr.Bool(true))
}

```
