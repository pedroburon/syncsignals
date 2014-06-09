package syncsignals

import (
    "container/list"
)

type callback func([]interface{}) error


type signal struct {
    Callbacks list.List
}

func (s *signal) connect(c callback) {
    s.Callbacks.PushBack(c)
}

func runCallback(ch chan error, fn callback, args []interface{}) {
    ch <- fn(args)
}

func (s *signal) send(args []interface{}) []error {
    errors := make([]error, s.Callbacks.Len())
    ch := make(chan error)
    for el:= s.Callbacks.Front(); el != nil; el = el.Next() {
        c := el.Value.(callback)
        go runCallback(ch, c, args)
    }
    for i := 0 ; i<s.Callbacks.Len() ; i++ {
        errors[i] = <-ch
    }
    return errors
}
