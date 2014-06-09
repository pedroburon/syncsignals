package syncsignals

import (
    "testing"
    "errors"
    "fmt"
    "github.com/bmizerany/assert"
)

func TestNew(t *testing.T) {
    s := new(signal)
    assert.Equal(t, 0, s.Callbacks.Len(), "Callback registered")
}

func TestConnect(t *testing.T) {
    s := new(signal)
    s.connect(func(args []interface{}) error {
        return nil
    })
    assert.Equal(t, 1, s.Callbacks.Len(), "Only one callback must be registered")
}

func TestSendOneCallback(t *testing.T) {
    s := new(signal)
    called := false
    s.connect(func(args []interface{}) error {
        called = true
        return nil
    })
    s.send(nil)
    assert.Equal(t, true, called, "Callback not called")
}

func TestSendOneCallbackTwice(t *testing.T) {
    s := new(signal)
    called := false
    s.connect(func(args []interface{}) error {
        called = true
        return nil
    })
    s.send(nil)
    called = false
    s.send(nil)
    assert.Equal(t, true, called, "1st Callback not called")
}

func TestSendTwoCallback(t *testing.T) {
    s := new(signal)
    called1st := false
    called2nd := false
    s.connect(func(args []interface{}) error {
        called1st = true
        return nil
    })
    s.connect(func(args []interface{}) error {
        called2nd = true
        return nil
    })
    s.send(nil)
    assert.Equal(t, true, called1st, "1st Callback not called")
    assert.Equal(t, true, called2nd, "2nd Callback not called")
}

func TestSendWithArgs(t *testing.T) {
    s := new(signal)
    calledArgs := []interface{}{}
    s.connect(func(args []interface{}) error {
        calledArgs = args
        return nil
    })
    expected := []interface{}{"hola", "mundo"}
    s.send(expected)
    assert.Equal(t, expected, calledArgs, "function not called with proper args")
}

func assertSameErrors(t *testing.T, expected []error, results []error) {
    assert.Equal(t, len(expected), len(results))
    for el := range expected {
        isThere := false
        for r := range results {
            if r == el {
                isThere = true
                break
            }
        }
        assert.Equal(t, true, isThere, fmt.Sprintf("element %s not present", el))
    }
}

func TestReturnTwoCallback(t *testing.T) {
    s := new(signal)
    err1 := errors.New("Error1")
    err2 := errors.New("Error2")
    s.connect(func(args []interface{}) error {
        return err1
    })
    s.connect(func(args []interface{}) error {
        return err2
    })
    results := s.send(nil)
    expected := []error{err1, err2}
    
    assertSameErrors(t, expected, results)
    
    // assert.Equal(t, results, expected, "results diferent from expected")
}

func TestReturnCallback(t *testing.T) {
    s := new(signal)
    err := errors.New("Mock error")
    s.connect(func(args []interface{}) error {
        return err
    })
    results := s.send(nil)
    expected := []error{err}

    assertSameErrors(t, expected, results)
}
