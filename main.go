package jsonNull

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Bool struct {
	sql.NullBool
}

func (v Bool) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Bool)
	} else {
		return json.Marshal(nil)
	}
}

func (v *Bool) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *bool
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Bool = *x
	} else {
		v.Valid = false
	}
	return nil
}

type Byte struct {
	sql.NullByte
}

func (v Byte) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Byte)
	} else {
		return json.Marshal(nil)
	}
}

func (v *Byte) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *byte
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Byte = *x
	} else {
		v.Valid = false
	}
	return nil
}

type DateTime struct {
	sql.NullTime
}

func (v DateTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (v *DateTime) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *time.Time
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Time = *x
	} else {
		v.Valid = false
	}
	return nil
}

type Float64 struct {
	sql.NullFloat64
}

func (v Float64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Float64)
	} else {
		return json.Marshal(nil)
	}
}

func (v *Float64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *float64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Float64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

type Int16 struct {
	sql.NullInt16
}

func (v Int16) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int16)
	} else {
		return json.Marshal(nil)
	}
}

func (v *Int16) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *int16
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int16 = *x
	} else {
		v.Valid = false
	}
	return nil
}

type Int32 struct {
	sql.NullInt32
}

func (v Int32) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int32)
	} else {
		return json.Marshal(nil)
	}
}

func (v *Int32) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *int32
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int32 = *x
	} else {
		v.Valid = false
	}
	return nil
}

type Int64 struct {
	sql.NullInt64
}

func (v Int64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (v *Int64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

type String struct {
	sql.NullString
}

func (v String) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *String) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}

type Json struct {
	sql.NullString
}

func (v Json) MarshalJSON() ([]byte, error) {
	if v.Valid {
		// unmarshal as map
		var m map[string]any
		err := json.Unmarshal([]byte(v.String), &m)
		if err != nil {
			return nil, err
		}
		// remarshal map to bytes
		return json.Marshal(&m)
	} else {
		return json.Marshal(nil)
	}
}

func (v *Json) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}
