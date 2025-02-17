# Quiz: Effective Error Handling in Go

### Which of the following accurately describes how errors are returned and treated in Go?

- [ ] Errors are displayed as pop-up messages.
- [x] Errors are returned as values from functions.
- [ ] Errors are automatically corrected by the runtime.
- [ ] Errors are stored in a global variable.

> [!NOTE]
> Errors are returned as values from functions meaning that they can be handled similarly to any other datatype.

### Which of the following best describes the purpose of the built-in error type in Go?

- [ ] Check for error conditions and return true or false.
- [ ] Store error codes and messages in a struct.
- [ ] Define common error messages used in Go programs.
- [x] Represent an error condition with a descriptive message.

> [!NOTE]
> The built-in error type in Go is an interface designed to represent an error condition with a descriptive message.

### In the provided code snippet, what is the purpose of creating a custom error type `CustomError`?

```go
type CustomError struct {
    Code    int
    Message string
}

func NewCustomError(code int, message string) *CustomError {
    return &CustomError{Code: code, Message: message}
}

func main() {
    err := NewCustomError(404, "Not Found")
    if err != nil {
        fmt.Println("Error Code:", err.Code, "Message:", err.Message)
    }
}
```

- [ ] To automatically log errors to a file
- [ ] To replace the built-in error type in Go
- [ ] To handle errors without using the error interface
- [x] To embed additional context or information within error messages

> [!NOTE]
> The CustomError struct allows developers to include additional information like error codes and messages to provide more context about the error.

### Complete the following code to demonstrate the use of panic() and recover() functions:

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    startApplication()
    fmt.Println("Application started successfully")
}

func startApplication() {
    if _, err := os.Stat("app.config"); os.IsNotExist(err) {
        panic("Application configuration file is missing")
    }
    fmt.Println("Configuration loaded successfully")
}
```

> [!NOTE]
>  `panic()` prints a panic message and begins unwinding the stack, then runs deferred functions and terminates the program. `recover()` is used to recover from a panic and resume normal execution.

### In the provided code snippet, what does the %w in fmt.Errorf do?

```go
func readFile(filename string) error {
    originalErr := errors.New("file not found")
    return fmt.Errorf("failed to read file: %w", originalErr)
}
```
- [ ] It prints the error message in uppercase.
- [ ] It logs the error to a file.
- [x] It wraps the original error with additional context.
- [ ] It converts the error message to a warning.

> [!NOTE]
> The `%w` in `fmt.Errorf` is used to wrap the original error with more context, preserving the original error information.
