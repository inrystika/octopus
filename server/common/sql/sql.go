package sql

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type SqlJson map[string]string

func (r SqlJson) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *SqlJson) Scan(input interface{}) error {
	switch v := input.(type) {
	case []byte:
		return json.Unmarshal(input.([]byte), r)
	default:
		return fmt.Errorf("cannot Scan() from: %#v", v)
	}
}
