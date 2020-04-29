// Package entities contains generated structures.
package entities



// RequiredNullable structure is generated from "#".
type RequiredNullable struct {
	A *Def    `json:"a"` // Required.
	B *string `json:"b"` // Required.
}

// Def structure is generated from "#/definitions/def".
type Def struct {
	B string `json:"b,omitempty"`
}
