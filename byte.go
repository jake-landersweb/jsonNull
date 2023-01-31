package jsonNull

import (
	"database/sql"
	"encoding/json"
)

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
