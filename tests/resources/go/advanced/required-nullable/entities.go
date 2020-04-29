// Package entities contains generated structures.
package entities



// RequiredNullable structure is generated from "#".
type RequiredNullable struct {
	A   *Def          `json:"a"`   // Required.
	An  *DefNullable  `json:"an"`  // Required.
	Axn *DefXNullable `json:"axn"` // Required.
	B   *string       `json:"b"`   // Required.
}

// Def structure is generated from "#/definitions/def".
type Def struct {
	B string `json:"b,omitempty"`
}

// DefNullable structure is generated from "#/definitions/defNullable".
type DefNullable struct {
	B string `json:"b,omitempty"`
}

// DefXNullable structure is generated from "#/definitions/defXNullable".
type DefXNullable struct {
	B string `json:"b,omitempty"`
}
