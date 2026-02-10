package exercises

import (
	"context"
	"errors"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ErrEmptyName is returned when name is empty
var ErrEmptyName = errors.New("name cannot be empty")

// GreetInput is the input for the Greet tool
type GreetInput struct {
	Name string `json:"name"`
}

// GreetOutput is the output for the Greet tool
type GreetOutput struct {
	Greeting string `json:"greeting"`
}

// TODO: Define GreetTool
// - Name: "greet"
// - Description: "Greets a person by name"

var GreetTool *mcp.Tool // Replace with actual definition

// TODO: Implement GreetHandler function
// - If Name is empty, return (nil, ErrEmptyName)
// - Otherwise return greeting: "Hello, {Name}!"
//
// Example:
//   GreetHandler with Name="Alice" returns GreetOutput{Greeting: "Hello, Alice!"}
//   GreetHandler with Name="" returns error ErrEmptyName

func GreetHandler(ctx context.Context, sess *mcp.ServerSession, params *mcp.CallToolParamsFor[GreetInput]) (*mcp.CallToolResultFor[GreetOutput], error) {
	// TODO: Implement this function
	return nil, nil
}
