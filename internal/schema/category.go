package schema

import (
	"fmt"
)

type Category string

var categoryValues = []string{
	"poetry",
	"drama",
	"prose",
	"non-fiction",
	"media",
}

func (c *Category) UnmarshalText(text []byte) error {
	s := string(text)
	for i := range categoryValues {
		if categoryValues[i] == s {
			*c = Category(s)
			return nil
		}
	}
	return fmt.Errorf("invalid value %q", string(text))
}

func (c *Category) String() string {
	return string(*c)
}

var InvalidCategory = Category("")
