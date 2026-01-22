# Go - Level 2: Understand

Explain ideas and concepts in your own words.

## Questions

1. Explain why we use a separate `tools/` package instead of putting everything in `main`.

2. Describe what `CallToolParamsFor[PingInput]` means - what is the `[PingInput]` part?

3. Why does the `Ping` function return two values `(*CallToolResultFor[PingOutput], error)`?

4. Explain the difference between `PingInput struct{}` and `PingInput struct{ Name string }`.

5. Why do we use pointers (`*mcp.ServerSession`) instead of values in function parameters?

## Answers

<details>
<summary>Click to reveal answers</summary>

1. **Separation of concerns** - The `tools/` package handles tool logic while `main` handles server setup. This makes code more maintainable, testable, and allows tools to be reused or tested independently.

2. **Go generics** - `CallToolParamsFor` is a generic type and `[PingInput]` is the type parameter. It means "CallToolParams specifically for PingInput". The compiler ensures type safety at compile time.

3. **Go error handling convention** - Go functions that can fail return `(result, error)`. Callers check `if err != nil` to handle errors. The pointer `*CallToolResultFor` allows returning `nil` on error.

4. **Empty vs populated struct** - `struct{}` has zero fields and zero size (used when no input needed). `struct{ Name string }` has one field and carries data. Empty structs are memory-efficient markers.

5. **Efficiency and mutability** - Pointers avoid copying large structs and allow functions to see modifications. `*ServerSession` is a pointer because sessions carry state that may be modified.

</details>
