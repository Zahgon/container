package container

// MustSingleton wraps the `Singleton` method and panics on errors instead of returning the errors.
func MustSingleton(c Container, resolver interface{}) { _ = "STUB: not implemented"; return }

// MustSingleton wraps the `SingletonLazy` method and panics on errors instead of returning the errors.
func MustSingletonLazy(c Container, resolver interface{}) { _ = "STUB: not implemented"; return }

// MustNamedSingleton wraps the `NamedSingleton` method and panics on errors instead of returning the errors.
func MustNamedSingleton(c Container, name string, resolver interface{}) {
	_ = "STUB: not implemented"
	return
}

// MustNamedSingleton wraps the `NamedSingletonLazy` method and panics on errors instead of returning the errors.
func MustNamedSingletonLazy(c Container, name string, resolver interface{}) {
	_ = "STUB: not implemented"
	return
}

// MustTransient wraps the `Transient` method and panics on errors instead of returning the errors.
func MustTransient(c Container, resolver interface{}) { _ = "STUB: not implemented"; return }

// MustTransientLazy wraps the `TransientLazy` method and panics on errors instead of returning the errors.
func MustTransientLazy(c Container, resolver interface{}) { _ = "STUB: not implemented"; return }

// MustNamedTransient wraps the `NamedTransient` method and panics on errors instead of returning the errors.
func MustNamedTransient(c Container, name string, resolver interface{}) {
	_ = "STUB: not implemented"
	return
}

// MustNamedTransient wraps the `NamedTransientLazy` method and panics on errors instead of returning the errors.
func MustNamedTransientLazy(c Container, name string, resolver interface{}) {
	_ = "STUB: not implemented"
	return
}

// MustCall wraps the `Call` method and panics on errors instead of returning the errors.
func MustCall(c Container, receiver interface{}) { _ = "STUB: not implemented"; return }

// MustResolve wraps the `Resolve` method and panics on errors instead of returning the errors.
func MustResolve(c Container, abstraction interface{}) { _ = "STUB: not implemented"; return }

// MustNamedResolve wraps the `NamedResolve` method and panics on errors instead of returning the errors.
func MustNamedResolve(c Container, abstraction interface{}, name string) {
	_ = "STUB: not implemented"
	return
}

// MustFill wraps the `Fill` method and panics on errors instead of returning the errors.
func MustFill(c Container, receiver interface{}) { _ = "STUB: not implemented"; return }
