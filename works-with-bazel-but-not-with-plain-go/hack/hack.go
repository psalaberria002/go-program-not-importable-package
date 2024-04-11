package hack

// This is a hack :( for making go mod tidy not remove these dependencies from go.mod
// The dependency is used as a binary target (@).
import (
	_ "github.com/jimmidyson/configmap-reload"
)
