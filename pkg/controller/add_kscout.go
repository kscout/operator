package controller

import (
	"github.com/kscout/operator/pkg/controller/kscout"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, kscout.Add)
}
