package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/robertlopez/go-mcp/styleguide"
)

// GetUberStylesInput is the input schema for the get_uber_styles tool.
type GetUberStylesInput struct {
	Code string `json:"code"`
}

// StyleSection is a matched section returned by the tool.
type StyleSection struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
}

// GetUberStylesOutput is the structured output of the get_uber_styles tool.
type GetUberStylesOutput struct {
	Sections []StyleSection `json:"sections"`
	Count    int            `json:"count"`
	Version  string         `json:"version"`
}

// GetUberStylesTool is the MCP tool definition for get_uber_styles.
var GetUberStylesTool = &mcp.Tool{
	Name:        "get_uber_styles",
	Description: "Analyze Go code against the Uber Go Style Guide. Returns relevant style guide sections with full explanations and code examples. Pass the Go code you want to validate.",
	Annotations: &mcp.ToolAnnotations{
		ReadOnlyHint: true,
	},
}

// NewGetUberStylesHandler returns a handler closure that captures the loaded StyleGuide.
func NewGetUberStylesHandler(sg *styleguide.StyleGuide) func(ctx context.Context, sess *mcp.ServerSession, params *mcp.CallToolParamsFor[GetUberStylesInput]) (*mcp.CallToolResultFor[GetUberStylesOutput], error) {
	return func(ctx context.Context, sess *mcp.ServerSession, params *mcp.CallToolParamsFor[GetUberStylesInput]) (*mcp.CallToolResultFor[GetUberStylesOutput], error) {
		code := params.Arguments.Code
		if code == "" {
			return &mcp.CallToolResultFor[GetUberStylesOutput]{
				IsError: true,
				Content: []mcp.Content{
					&mcp.TextContent{Text: "code parameter is required and must be non-empty"},
				},
			}, nil
		}

		matched := sg.MatchSections(code)

		sections := make([]StyleSection, len(matched))
		for i, s := range matched {
			sections[i] = StyleSection{
				ID:       s.ID,
				Title:    s.Title,
				Category: s.Category,
				Content:  s.Content,
			}
		}

		output := GetUberStylesOutput{
			Sections: sections,
			Count:    len(sections),
			Version:  "e2c8a0e",
		}

		// Build text summary for Content array.
		summary := fmt.Sprintf("Found %d relevant style guide sections.", output.Count)
		if output.Count > 0 {
			summary += "\n\nSections:"
			for _, s := range sections {
				summary += fmt.Sprintf("\n- [%s] %s", s.Category, s.Title)
			}
		}

		// Marshal structured content for inclusion as text fallback.
		jsonBytes, _ := json.MarshalIndent(output, "", "  ")

		return &mcp.CallToolResultFor[GetUberStylesOutput]{
			Content: []mcp.Content{
				&mcp.TextContent{Text: summary},
				&mcp.TextContent{Text: string(jsonBytes)},
			},
			StructuredContent: output,
		}, nil
	}
}
