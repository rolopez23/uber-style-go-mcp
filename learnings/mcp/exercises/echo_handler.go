package exercises

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// EchoInput is the input for the Echo tool
type EchoInput struct {
	Message string `json:"message"`
}

// EchoOutput is the output for the Echo tool
type EchoOutput struct {
	Echo string `json:"echo"`
}

// TODO: Implement Echo function
// - Takes context, session, and params
// - Returns the input message in the Echo field of EchoOutput
//
// Signature:
//   func Echo(ctx context.Context, sess *mcp.ServerSession, params *mcp.CallToolParamsFor[EchoInput]) (*mcp.CallToolResultFor[EchoOutput], error)
//
// Example implementation pattern:
//   return &mcp.CallToolResultFor[EchoOutput]{
//       StructuredContent: EchoOutput{
//           Echo: params.Arguments.Message,
//       },
//   }, nil

func Echo(ctx context.Context, sess *mcp.ServerSession, params *mcp.CallToolParamsFor[EchoInput]) (*mcp.CallToolResultFor[EchoOutput], error) {
	// TODO: Implement this function
	return nil, nil
}
