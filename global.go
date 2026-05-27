package container

// Global is the global concrete of the Container.
var Global = New()

// Singleton calls the same method of the global concrete.
func Singleton(resolver interface{}) error { _ = "STUB: not implemented"; return nil }

// SingletonLazy calls the same method of the global concrete.
func SingletonLazy(resolver interface{}) error { _ = "STUB: not implemented"; return nil }

// NamedSingleton calls the same method of the global concrete.
func NamedSingleton(name string, resolver interface{}) error { _ = "STUB: not implemented"; return nil }

// NamedSingletonLazy calls the same method of the global concrete.
func NamedSingletonLazy(name string, resolver interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transient calls the same method of the global concrete.
func Transient(resolver interface{}) error { _ = "STUB: not implemented"; return nil }

// TransientLazy calls the same method of the global concrete.
func TransientLazy(resolver interface{}) error { _ = "STUB: not implemented"; return nil }

// NamedTransient calls the same method of the global concrete.
func NamedTransient(name string, resolver interface{}) error { _ = "STUB: not implemented"; return nil }

// NamedTransientLazy calls the same method of the global concrete.
func NamedTransientLazy(name string, resolver interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Reset calls the same method of the global concrete.
func Reset() {
	_ = "STUB: not implemented"

	// Call calls the same method of the global concrete.
	return
}

func Call(receiver interface{}) error { _ = "STUB: not implemented"; return nil }

// Resolve calls the same method of the global concrete.
func Resolve(abstraction interface{}) error { _ = "STUB: not implemented"; return nil }

// NamedResolve calls the same method of the global concrete.
func NamedResolve(abstraction interface{}, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// Fill calls the same method of the global concrete.
func Fill(receiver interface{}) error { _ = "STUB: not implemented"; return nil }
