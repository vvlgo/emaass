package ecode

import (
	"fmt"
	"strconv"
	"sync/atomic"
)

var (
	_messages atomic.Value         // NOTE: stored map[string]map[int]string
	_codes    = map[int]struct{}{} // register codes.
)

// Register register ecode message map.
func register(cm map[int]string) {
	_messages.Store(cm)
}

// New new a ecode.Codes by int value.
// NOTE: ecode must unique in global, the New will check repeat and then panic.
func New(e int, message string) Code {
	if e < 0 {
		panic("business ecode must greater than zero")
	}
	return add(e, message)
}

func add(e int, message string) Code {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	m := _messages.Load()
	if m == nil {
		register(map[int]string{e: message})
	} else {
		m := m.(map[int]string)
		m[e] = message
		register(m)
	}

	return Int(e)
}

// Codes ecode error interface which has a code & message.
type Codes interface {
	// sometimes Error return Code in string form
	// NOTE: don't use Error in monitor report even it also work for now
	Error() string
	// Code get error code.
	Code() int
	// Message get code message.
	Message() string
}

// A Code is an int error code spec.
type Code int

func (e Code) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

// Code return error code
func (e Code) Code() int { return int(e) }

// Message return error message
func (e Code) Message() string {
	if cm, ok := _messages.Load().(map[int]string); ok {
		if msg, ok := cm[e.Code()]; ok {
			return msg
		}
	}
	return e.Error()
}

// Details return details.
func (e Code) Details() []interface{} { return nil }

// Int parse code int to error.
func Int(i int) Code { return Code(i) }
