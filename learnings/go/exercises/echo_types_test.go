// Run: go test ./learnings/go/exercises/... -run TestEchoInput -v

package exercises

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestEchoInput_HasMessageField(t *testing.T) {
	input := EchoInput{Message: "hello"}
	if input.Message != "hello" {
		t.Errorf("expected Message to be 'hello', got %q", input.Message)
	}
}

func TestEchoInput_JSONTag(t *testing.T) {
	typ := reflect.TypeOf(EchoInput{})
	field, ok := typ.FieldByName("Message")
	if !ok {
		t.Fatal("EchoInput must have a Message field")
	}

	tag := field.Tag.Get("json")
	if tag != "message" {
		t.Errorf("expected json tag 'message', got %q", tag)
	}
}

func TestEchoInput_JSONMarshal(t *testing.T) {
	input := EchoInput{Message: "test"}
	data, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	expected := `{"message":"test"}`
	if string(data) != expected {
		t.Errorf("expected %s, got %s", expected, string(data))
	}
}
