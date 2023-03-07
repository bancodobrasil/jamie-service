package dtos

import v1 "github.com/bancodobrasil/jamie-service/payloads/v1"

// Eval ...
type Eval map[string]interface{}

// NewEval ...
func NewEval(payloads v1.Eval) Eval {
	return Eval(payloads)
}
