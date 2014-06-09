syncsignals
===========

Synchronous Signals Django like for golang.


Usage
-----

    package main

    import (
        "github.com/pedroburon/syncsignals.git"
        "fmt"
    )

    postSomething = new(signal)

    func doSomething() {
        // do something
        args := []interface{}{"hola", "mundo"}
        postSomething.send(args)
    }

    func main() {
        postSomething.register(func(args []interface{}) error{
            // do something after do something
            fmt.Println(args) // Will print "[hola mundo]"
            return nil // no error
        })
        doSomething()// do something then...
    }
