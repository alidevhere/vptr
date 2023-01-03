# What is conv ? 
This package provides a set of functions for converting values to pointers.

There are scenarios when you have to use a pointer to a value without cerating a variable. For example





Most developers use this syntax to convert value to reference to value:
&[]type{}[0].
But it creates the code ugly


# Usage
It does not convert types, it just converts values to pointers.
It is useful for converting values to pointers in the function parameters.
For example:

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

It is possible to use this package to convert values to pointers:
func f(a *int, b *string, c *bool) {
	...
}

func main() {
	f(Int(1), String("2"), Bool(true))
}

```