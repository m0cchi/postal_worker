package module

import (
	"github.com/m0cchi/postal_worker/lib/model"
)

// Module Interface
type PostalModule interface {
	GetModuleName() string
	Exec(message model.PostalMatter, to model.To) error
}
