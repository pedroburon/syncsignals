package syncsignals

import (
    "container/list"
    "fmt"
)

type callback func([]interface{}) error


type signal struct {
    Callbacks list.List
}

func (s *signal) connect(c callback) {
    s.Callbacks.PushBack(c)
}

func (s *signal) send(args []interface{}) []error {
    errors := make([]error, s.Callbacks.Len())
    index := 0
    for el:= s.Callbacks.Front(); el != nil; el = el.Next() {
        callback := el.Value.(callback)
        errors[index] = callback(args)
        index += 1
    }
    return errors
}
