# go-exercises
Some random exercises that I made in Go

## Useful powershell commands

* go env {variable}: gets the value of a GO environment variable.

    for example, to get the variable "GO operating System":
```go
    go env GOOS
```

* go env -w {variable}="{value}": writes the value of a GO enviroment variable.

    for example, to change the variable "GO operating System" to "linux":
```go
    go env -w GOOS="linux"
```


* go env -u {variable}

    for example, to change back the variable "GO operating System" to the default value:
```go
    go env -u GOOS
```

* To build and an application (generating a .exe app):

```go
    go build .
```

* To run an application without build:

```go
    go run .
```

## Useful libraries

* fmt: Format library
* math: Math library
* date: Date and Times library
* sort: To Sort results

## Language

* to initialize a module

```go
    // In the current folder of the main file, execute this:
    go mod init jgqsolutions.com.ar/structs    
```

* to create a new file:

```go
    // 1) add a new .go file
    // 2) In the first line, add package main (or the name of the module)
    package main
    // 3) after adding the main package, you need to add a main function
    func main() {
        
    }
```

* to import libraries:

```go
    import {
        "fmt"
    }
```

* to define an array with fixed elements:
```go
    var colors [3]string
    colors[0] = "Red"
    colors[1] = "Blue"
    colors[2] = "Green"
```

another way:

```go
numbers = make([]string, 3, 3)
```

* to define an array with fixed elements inline:
```go
    var colors = [3]string {"Red", "Blue", "Green"}
```

* to define an array with variable elements (slice):
```go
    var colors []string
    colors[0] = "Red"
    colors[1] = "Blue"
    colors[2] = "Green"
```

another way:

```go
numbers = make([]string, 3, 3)
```

* to define an array with variable elements inline (Slice):
```go
    var colors = []string {"Red", "Blue", "Green"}
```

* to define a dictionary:
```go
    states := make(map[string]string)
    fmt.Println(states)
    states["BA"] = "Buenos Aires"
    states["CB"] = "Córdoba"
    fmt.Println(states)
```

* switch break doesn't exist. Instead, if we want to continue to the next case segment, we need to use "fallthrough"

```go

    a := 1
    var result string
    switch a {
    case 1:
        result = "One"
        fallthrough
    case 2:
        result = "Two"
        fallthrough
    default:
        result = "Default"
    }

    // The result is Default
```

* Adding a new module

It's the same. In the calling module, if we need to test an internal module, we have to add this line to replace:

```go
go mod edit -replace {modulename}=../{module-folder}
```

Then, we need to call this line to update the references:

```go
go mod tidy
```

