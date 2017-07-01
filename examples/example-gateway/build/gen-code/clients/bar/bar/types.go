// Code generated by thriftrw v1.3.0
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type BarException struct {
	StringField string `json:"stringField,required"`
}

func (v *BarException) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.StringField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *BarException) FromWire(w wire.Value) error {
	var err error
	stringFieldIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.StringField, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				stringFieldIsSet = true
			}
		}
	}
	if !stringFieldIsSet {
		return errors.New("field StringField of BarException is required")
	}
	return nil
}

func (v *BarException) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("StringField: %v", v.StringField)
	i++
	return fmt.Sprintf("BarException{%v}", strings.Join(fields[:i], ", "))
}

func (v *BarException) Equals(rhs *BarException) bool {
	if !(v.StringField == rhs.StringField) {
		return false
	}
	return true
}

func (v *BarException) Error() string {
	return v.String()
}

type BarRequest struct {
	StringField string `json:"stringField,required"`
	BoolField   bool   `json:"boolField,required"`
}

func (v *BarRequest) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.StringField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueBool(v.BoolField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *BarRequest) FromWire(w wire.Value) error {
	var err error
	stringFieldIsSet := false
	boolFieldIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.StringField, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				stringFieldIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBool {
				v.BoolField, err = field.Value.GetBool(), error(nil)
				if err != nil {
					return err
				}
				boolFieldIsSet = true
			}
		}
	}
	if !stringFieldIsSet {
		return errors.New("field StringField of BarRequest is required")
	}
	if !boolFieldIsSet {
		return errors.New("field BoolField of BarRequest is required")
	}
	return nil
}

func (v *BarRequest) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("StringField: %v", v.StringField)
	i++
	fields[i] = fmt.Sprintf("BoolField: %v", v.BoolField)
	i++
	return fmt.Sprintf("BarRequest{%v}", strings.Join(fields[:i], ", "))
}

func (v *BarRequest) Equals(rhs *BarRequest) bool {
	if !(v.StringField == rhs.StringField) {
		return false
	}
	if !(v.BoolField == rhs.BoolField) {
		return false
	}
	return true
}

type BarResponse struct {
	StringField        string           `json:"stringField,required"`
	IntWithRange       int32            `json:"intWithRange,required"`
	IntWithoutRange    int32            `json:"intWithoutRange,required"`
	MapIntWithRange    map[string]int32 `json:"mapIntWithRange,required"`
	MapIntWithoutRange map[string]int32 `json:"mapIntWithoutRange,required"`
}

type _Map_String_I32_MapItemList map[string]int32

func (m _Map_String_I32_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		kw, err := wire.NewValueString(k), error(nil)
		if err != nil {
			return err
		}
		vw, err := wire.NewValueI32(v), error(nil)
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_String_I32_MapItemList) Size() int {
	return len(m)
}

func (_Map_String_I32_MapItemList) KeyType() wire.Type {
	return wire.TBinary
}

func (_Map_String_I32_MapItemList) ValueType() wire.Type {
	return wire.TI32
}

func (_Map_String_I32_MapItemList) Close() {
}

func (v *BarResponse) ToWire() (wire.Value, error) {
	var (
		fields [5]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.StringField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueI32(v.IntWithRange), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	w, err = wire.NewValueI32(v.IntWithoutRange), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 3, Value: w}
	i++
	if v.MapIntWithRange == nil {
		return w, errors.New("field MapIntWithRange of BarResponse is required")
	}
	w, err = wire.NewValueMap(_Map_String_I32_MapItemList(v.MapIntWithRange)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 4, Value: w}
	i++
	if v.MapIntWithoutRange == nil {
		return w, errors.New("field MapIntWithoutRange of BarResponse is required")
	}
	w, err = wire.NewValueMap(_Map_String_I32_MapItemList(v.MapIntWithoutRange)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 5, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Map_String_I32_Read(m wire.MapItemList) (map[string]int32, error) {
	if m.KeyType() != wire.TBinary {
		return nil, nil
	}
	if m.ValueType() != wire.TI32 {
		return nil, nil
	}
	o := make(map[string]int32, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetString(), error(nil)
		if err != nil {
			return err
		}
		v, err := x.Value.GetI32(), error(nil)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

func (v *BarResponse) FromWire(w wire.Value) error {
	var err error
	stringFieldIsSet := false
	intWithRangeIsSet := false
	intWithoutRangeIsSet := false
	mapIntWithRangeIsSet := false
	mapIntWithoutRangeIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.StringField, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				stringFieldIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TI32 {
				v.IntWithRange, err = field.Value.GetI32(), error(nil)
				if err != nil {
					return err
				}
				intWithRangeIsSet = true
			}
		case 3:
			if field.Value.Type() == wire.TI32 {
				v.IntWithoutRange, err = field.Value.GetI32(), error(nil)
				if err != nil {
					return err
				}
				intWithoutRangeIsSet = true
			}
		case 4:
			if field.Value.Type() == wire.TMap {
				v.MapIntWithRange, err = _Map_String_I32_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
				mapIntWithRangeIsSet = true
			}
		case 5:
			if field.Value.Type() == wire.TMap {
				v.MapIntWithoutRange, err = _Map_String_I32_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
				mapIntWithoutRangeIsSet = true
			}
		}
	}
	if !stringFieldIsSet {
		return errors.New("field StringField of BarResponse is required")
	}
	if !intWithRangeIsSet {
		return errors.New("field IntWithRange of BarResponse is required")
	}
	if !intWithoutRangeIsSet {
		return errors.New("field IntWithoutRange of BarResponse is required")
	}
	if !mapIntWithRangeIsSet {
		return errors.New("field MapIntWithRange of BarResponse is required")
	}
	if !mapIntWithoutRangeIsSet {
		return errors.New("field MapIntWithoutRange of BarResponse is required")
	}
	return nil
}

func (v *BarResponse) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [5]string
	i := 0
	fields[i] = fmt.Sprintf("StringField: %v", v.StringField)
	i++
	fields[i] = fmt.Sprintf("IntWithRange: %v", v.IntWithRange)
	i++
	fields[i] = fmt.Sprintf("IntWithoutRange: %v", v.IntWithoutRange)
	i++
	fields[i] = fmt.Sprintf("MapIntWithRange: %v", v.MapIntWithRange)
	i++
	fields[i] = fmt.Sprintf("MapIntWithoutRange: %v", v.MapIntWithoutRange)
	i++
	return fmt.Sprintf("BarResponse{%v}", strings.Join(fields[:i], ", "))
}

func _Map_String_I32_Equals(lhs, rhs map[string]int32) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for lk, lv := range lhs {
		rv, ok := rhs[lk]
		if !ok {
			return false
		}
		if !(lv == rv) {
			return false
		}
	}
	return true
}

func (v *BarResponse) Equals(rhs *BarResponse) bool {
	if !(v.StringField == rhs.StringField) {
		return false
	}
	if !(v.IntWithRange == rhs.IntWithRange) {
		return false
	}
	if !(v.IntWithoutRange == rhs.IntWithoutRange) {
		return false
	}
	if !_Map_String_I32_Equals(v.MapIntWithRange, rhs.MapIntWithRange) {
		return false
	}
	if !_Map_String_I32_Equals(v.MapIntWithoutRange, rhs.MapIntWithoutRange) {
		return false
	}
	return true
}

type UUID string

func (v UUID) ToWire() (wire.Value, error) {
	x := (string)(v)
	return wire.NewValueString(x), error(nil)
}

func (v UUID) String() string {
	x := (string)(v)
	return fmt.Sprint(x)
}

func (v *UUID) FromWire(w wire.Value) error {
	x, err := w.GetString(), error(nil)
	*v = (UUID)(x)
	return err
}

func (lhs UUID) Equals(rhs UUID) bool {
	return (lhs == rhs)
}
