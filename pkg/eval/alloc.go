// Copyright 2016 Marapongo, Inc. All rights reserved.

package eval

import (
	"github.com/marapongo/mu/pkg/compiler/symbols"
	"github.com/marapongo/mu/pkg/diag"
	"github.com/marapongo/mu/pkg/eval/rt"
)

// Allocator is a factory for creating objects.
type Allocator struct {
	hooks InterpreterHooks // an optional set of allocation lifetime callback hooks.
}

// NewAllocator allocates a fresh allocator instance.
func NewAllocator(hooks InterpreterHooks) *Allocator {
	return &Allocator{hooks: hooks}
}

// onNewObject is invoked for each allocation and emits an appropriate event.
func (a *Allocator) onNewObject(o *rt.Object) {
	if a.hooks != nil {
		a.hooks.OnNewObject(o)
	}
}

// New creates a new empty object of the given type.
func (a *Allocator) New(t symbols.Type, properties rt.PropertyMap) *rt.Object {
	obj := rt.NewObject(t, nil, properties)
	a.onNewObject(obj)
	return obj
}

// NewArray creates a new array object of the given element type.
func (a *Allocator) NewArray(elem symbols.Type, arr *[]*rt.Pointer) *rt.Object {
	obj := rt.NewArrayObject(elem, arr)
	a.onNewObject(obj)
	return obj
}

// NewPrimitive creates a new primitive object with the given primitive type.
func (a *Allocator) NewPrimitive(t symbols.Type, v interface{}) *rt.Object {
	obj := rt.NewPrimitiveObject(t, v)
	a.onNewObject(obj)
	return obj
}

// NewBool creates a new primitive number object.
func (a *Allocator) NewBool(v bool) *rt.Object {
	obj := rt.NewBoolObject(v)
	a.onNewObject(obj)
	return obj
}

// NewNumber creates a new primitive number object.
func (a *Allocator) NewNumber(v float64) *rt.Object {
	obj := rt.NewNumberObject(v)
	a.onNewObject(obj)
	return obj
}

// NewNull creates a new null object.
func (a *Allocator) NewNull() *rt.Object {
	obj := rt.NewNullObject()
	a.onNewObject(obj)
	return obj
}

// NewString creates a new primitive number object.
func (a *Allocator) NewString(v string) *rt.Object {
	obj := rt.NewStringObject(v)
	a.onNewObject(obj)
	return obj
}

// NewFunction creates a new function object that can be invoked, with the given symbol.
func (a *Allocator) NewFunction(fnc symbols.Function, this *rt.Object) *rt.Object {
	obj := rt.NewFunctionObject(fnc, this)
	a.onNewObject(obj)
	return obj
}

// NewPointer allocates a new pointer-like object that wraps the given reference.
func (a *Allocator) NewPointer(t symbols.Type, ptr *rt.Pointer) *rt.Object {
	obj := rt.NewPointerObject(t, ptr)
	a.onNewObject(obj)
	return obj
}

// NewException creates a new exception with the given message.
func (a *Allocator) NewException(node diag.Diagable, stack *rt.StackFrame,
	message string, args ...interface{}) *rt.Object {
	obj := rt.NewExceptionObject(node, stack, message, args...)
	a.onNewObject(obj)
	return obj
}

// NewConstant returns a new object with the right type and value, based on some constant data.
func (a *Allocator) NewConstant(v interface{}) *rt.Object {
	obj := rt.NewConstantObject(v)
	a.onNewObject(obj)
	return obj
}
