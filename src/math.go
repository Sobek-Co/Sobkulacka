package main

//#include "ourmath.h"
import "C"

import (
	"fmt"
	"runtime"
)

type Error int

const (
	Success        Error = C.SUCESS
	DivisionByZero Error = C.DIVISION_BY_ZERO
)

type State struct {
	ptr *C.m_state
}

func NewState() *State {
	s := &State{
		ptr: C.m_init_state(),
	}
	runtime.SetFinalizer(s, func(s *State) {
		s.Free()
	})
	return s
}

func (s *State) Check() error {
	switch s.Error() {
	case Success:
		return nil
	case DivisionByZero:
		return fmt.Errorf("division by zero")
	default:
		return fmt.Errorf("unknown error")
	}
}

func (s *State) Free() {
	if s.ptr != nil {
		C.m_free_state(s.ptr)
		s.ptr = nil
	}
}

func (s *State) Clear() {
	C.m_clear(s.ptr)
}

func (s *State) Value() float64 {
	return float64(C.m_get_state(s.ptr))
}

func (s *State) Error() Error {
	return Error(C.m_get_err(s.ptr))
}

func (s *State) Set(value float64) *State {
	C.m_set_state(s.ptr, C.double(value))
	return s
}

func (s *State) Add(v float64) *State {
	C.m_add(s.ptr, C.double(v))
	return s
}

func (s *State) Sub(v float64) *State {
	C.m_sub(s.ptr, C.double(v))
	return s
}

func (s *State) Mul(v float64) *State {
	C.m_mul(s.ptr, C.double(v))
	return s
}

func (s *State) Div(v float64) *State {
	C.m_div(s.ptr, C.double(v))
	return s
}

func (s *State) Mod(v int) *State {
	C.m_mod(s.ptr, C.int(v))
	return s
}

func (s *State) Power(v int) *State {
	C.m_power(s.ptr, C.int(v))
	return s
}

func (s *State) Root(v int) *State {
	C.m_root(s.ptr, C.int(v))
	return s
}

func (s *State) Factorial() *State {
	C.m_factorial(s.ptr)
	return s
}

func (s *State) Sin() *State {
	C.m_sin(s.ptr)
	return s
}

func (s *State) Cos() *State {
	C.m_cos(s.ptr)
	return s
}

func (s *State) Tan() *State {
	C.m_tg(s.ptr)
	return s
}

func (s *State) Cot() *State {
	C.m_cotg(s.ptr)
	return s
}
