package jsonNull

import (
	"database/sql"
	"encoding/json"
)

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
