package metrics

import (
	"reflect"
	"strings"

	"github.com/aopoltorzhicky/bcdhub/internal/models"
)

// Bool -
type Bool struct {
	Field string
}

// NewBool -
func NewBool(field string) *Bool {
	return &Bool{
		Field: field,
	}
}

// Compute -
func (m *Bool) Compute(a, b models.Contract) Feature {
	f := Feature{
		Name: strings.ToLower(m.Field),
	}
	aVal := m.getContractField(a)
	bVal := m.getContractField(b)

	if aVal == bVal {
		f.Value = 1
	}
	return f
}

func (m *Bool) getContractField(c models.Contract) interface{} {
	r := reflect.ValueOf(c)
	return reflect.Indirect(r).FieldByName(m.Field).Interface()
}
