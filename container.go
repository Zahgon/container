// Package container is a lightweight yet powerful IoC container for Go projects.
// It provides an easy-to-use interface and performance-in-mind container to be your ultimate requirement.
package container

import (
	"reflect"
)

// binding holds a resolver and a concrete (if already resolved).
// It is the break for the Container wall!
type binding struct {
	resolver    interface{} // resolver is the function that is responsible for making the concrete.
	concrete    interface{} // concrete is the stored instance for singleton bindings.
	isSingleton bool        // isSingleton is true if the binding is a singleton.
}

// make resolves the binding if needed and returns the resolved concrete.
func (b *binding) make(c Container) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Container holds the bindings and provides methods to interact with them.
// It is the entry point in the package.
type Container map[reflect.Type]map[string]*binding

// New creates a new concrete of the Container.
func New() Container {
	_ = "STUB: not implemented"
	return *

	// bind maps an abstraction to concrete and instantiates if it is a singleton binding.
	new(Container)
}

func (c Container) bind(resolver interface{}, name string, isSingleton bool, isLazy bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c Container) validateResolverFunction(funcType reflect.Type) error {
	_ = "STUB: not implemented"
	return nil
}

// invoke calls a function and its returned values.
// It only accepts one value and an optional error.
func (c Container) invoke(function interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// arguments returns the list of resolved arguments for a function.
func (c Container) arguments(function interface{}) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reset deletes all the existing bindings and empties the container.
func (c Container) Reset() { _ = "STUB: not implemented"; return }

// Singleton binds an abstraction to concrete in singleton mode.
// It takes a resolver function that returns the concrete, and its return type matches the abstraction (interface).
// The resolver function can have arguments of abstraction that have been declared in the Container already.
func (c Container) Singleton(resolver interface{}) error { _ = "STUB: not implemented"; return nil }

// SingletonLazy binds an abstraction to concrete lazily in singleton mode.
// The concrete is resolved only when the abstraction is resolved for the first time.
// It takes a resolver function that returns the concrete, and its return type matches the abstraction (interface).
// The resolver function can have arguments of abstraction that have been declared in the Container already.
func (c Container) SingletonLazy(resolver interface{}) error { _ = "STUB: not implemented"; return nil }

// NamedSingleton binds a named abstraction to concrete in singleton mode.
func (c Container) NamedSingleton(name string, resolver interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// NamedSingleton binds a named abstraction to concrete lazily in singleton mode.
// The concrete is resolved only when the abstraction is resolved for the first time.
func (c Container) NamedSingletonLazy(name string, resolver interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transient binds an abstraction to concrete in transient mode.
// It takes a resolver function that returns the concrete, and its return type matches the abstraction (interface).
// The resolver function can have arguments of abstraction that have been declared in the Container already.
func (c Container) Transient(resolver interface{}) error { _ = "STUB: not implemented"; return nil }

// TransientLazy binds an abstraction to concrete lazily in transient mode.
// Normally the resolver will be called during registration, but that is skipped in lazy mode.
// It takes a resolver function that returns the concrete, and its return type matches the abstraction (interface).
// The resolver function can have arguments of abstraction that have been declared in the Container already.
func (c Container) TransientLazy(resolver interface{}) error { _ = "STUB: not implemented"; return nil }

// NamedTransient binds a named abstraction to concrete lazily in transient mode.
func (c Container) NamedTransient(name string, resolver interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// NamedTransient binds a named abstraction to concrete in transient mode.
// Normally the resolver will be called during registration, but that is skipped in lazy mode.
func (c Container) NamedTransientLazy(name string, resolver interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Call takes a receiver function with one or more arguments of the abstractions (interfaces).
// It invokes the receiver function and passes the related concretes.
func (c Container) Call(function interface{}) error { _ = "STUB: not implemented"; return nil }

// Resolve takes an abstraction (reference of an interface type) and fills it with the related concrete.
func (c Container) Resolve(abstraction interface{}) error { _ = "STUB: not implemented"; return nil }

// NamedResolve takes abstraction and its name and fills it with the related concrete.
func (c Container) NamedResolve(abstraction interface{}, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// Fill takes a struct and resolves the fields with the tag `container:"inject"`
func (c Container) Fill(structure interface{}) error { _ = "STUB: not implemented"; return nil }
