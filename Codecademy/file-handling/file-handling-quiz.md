# File Handling Quiz

### What is the default behavior for performing os.Create(name) on an existing file?

Using os.Create(name) on an existing file will create a new file with a different number, such as file_2.
Using os.Create(name) on an existing file will do nothing to the existing file.
- [x] Using os.Create(name) on an existing file will truncate it, clearing that file’s contents.
- [ ] Using os.Create(name) on an existing file will cause an error because you cannot create a file that exists.

### Fill in the blanks to change the working directory in Go:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.Chdir("child_directory")
    if err != nil {
        fmt.Println("Error changing directory:", err)
        return
    }
}
```

### What types of operations can be performed using Go’s os package?

- [x] file handling, directory manipulation, environment variable management, and process control
network communication, socket programming, and data encryption
graphic rendering, animation, and user interface design
database querying, data modeling, and transaction management

### Fill in the blanks to create a directory in Go:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.Mkdir("src", 0777)
    if err != nil {
        fmt.Println("Error creating directory:", err)
        return
    }
}
```

### What is the correct way to close a file in Go to ensure it closes even if an error occurs?

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Open("my_book.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    /* What line of code goes here to properly close the file? */
}
```

- [ ] `f.Close()`
- [x] `defer f.Close()`
- [ ] `os.Close()`

### In Go, which `os.File` type methods are for writing data and strings to files?

- [ ] `File.Append(b) and File.AppendString(s)`
- [ ] `File.Read(b) and File.Write(s)`
- [x] `File.Write(b) and File.WriteString(s)`

### Fill in the blanks to read a file’s contents using Go:

```go
content, err := os.ReadFile("data.txt")
if err != nil {
    fmt.Println("Error reading file:", err)
    return
}

fmt.Println(string(content))
```

### Which of the following are NOT tasks involved in file handling within a computer’s filesystem?

- [ ] downloading, uploading, archiving, and compressing files
- [ ] reading, writing, creating, and deleting files
- [x] configuring network settings, monitoring system performance, and managing user accounts
- [ ] formatting, partitioning, encrypting, and backing up files