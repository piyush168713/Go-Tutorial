package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

/*
func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
}
*/

/*
To silence complaints about the unused imports, use a blank identifier to refer to a symbol from the 
imported package. Similarly, assigning the unused variable fd to the blank identifier will silence the 
unused variable error. This version of the program does compile.
*/


var _ = fmt.Printf // For debugging; delete when done.
var _ io.Reader    // For debugging; delete when done.

func main() {
    fd, err := os.Open("2_var.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
    _ = fd
}


/*
Import for side effect
An unused import like fmt or io in the previous example should eventually be used or removed: blank assignments 
identify code as a work in progress. But sometimes it is useful to import a package only for its side effects, 
without any explicit use. For example, during its init function, the net/http/pprof package registers 
HTTP handlers that provide debugging information. It has an exported API, but most clients need only the handler 
registration and access the data through a web page. To import the package only for its side effects, rename 
the package to the blank identifier:
*/

// import _ "net/http/pprof"