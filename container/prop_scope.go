package container

const (
	ScopePrototype = "prototype"
)

type defScope string

func (s defScope) isContainer() bool {
	return string(s) != ScopePrototype
}

func (s defScope) isPrototype() bool {
	return string(s) == ScopePrototype
}
