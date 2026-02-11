// Run: go test ./learnings/uber-style-guide/exercises/... -run TestClosureHandler -v

package exercises

// TODO: Implement NewGreetHandler
// Create a function that returns a handler closure capturing a greeting prefix.
//
// The returned function takes a name (string) and returns a greeting (string)
// by combining the captured prefix with the name.
//
// This demonstrates the same pattern used by NewGetUberStylesHandler:
// a constructor function that captures a dependency and returns a closure.
//
// Hints:
// - The outer function takes a prefix string parameter
// - The inner function (closure) takes a name string and returns a string
// - The closure captures the prefix from the outer function's scope
// - Use fmt.Sprintf or string concatenation to combine prefix + name
//
// Example:
//   handler := NewGreetHandler("Hello")
//   result := handler("World")  // "Hello, World!"
//
//   handler2 := NewGreetHandler("Hi")
//   result2 := handler2("Go")   // "Hi, Go!"

func NewGreetHandler(prefix string) func(name string) string {
	// TODO: Implement this function
	// Return a closure that captures prefix and uses it with name
	return nil
}
