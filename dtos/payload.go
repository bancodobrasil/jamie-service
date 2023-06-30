package dtos

import v1 "github.com/bancodobrasil/jamie-service/payloads/v1"

// Process ...
type Process map[string]interface{}

// NewEval ...
func NewProcess(payloads v1.Eval) Process {
	return Process(payloads)
}
