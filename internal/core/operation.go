package core

// Operation defines a single, serializable processing step.
type Operation struct {
	Name   string            `json:"name"`
	Params map[string]string `json:"params"`
}
