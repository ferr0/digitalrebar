package client

// Deprecated: use api instead. client will not be updated

import (
	"github.com/digitalrebar/digitalrebar/go/rebar-api/datatypes"
)

type Capability struct {
	datatypes.Capability
	Timestamps
	apiHelper
}
