package schema

import (
	"encoding"
)

var _ = interface {
	encoding.TextUnmarshaler
}(&InvalidCategory)
