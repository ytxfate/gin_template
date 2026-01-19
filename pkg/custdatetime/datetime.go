package custdatetime

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	dtDefaultFmt = time.DateTime
	dtCompactFmt = "20060102150405"
)

type DateTime time.Time

func (d *DateTime) UnmarshalJSON(data []byte) (err error) {
	ds := strings.Trim(string(data), "\"")
	t, err := time.Parse(dtDefaultFmt, ds)
	if err != nil {
		t, err = time.Parse(dtCompactFmt, ds)
		if err != nil {
			return err
		}
	}
	*d = DateTime(t)
	return
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", time.Time(d).Format(dtDefaultFmt))), nil
}

// Implement driver.Valuer for saving JSONB
func (d DateTime) Value() (driver.Value, error) {
	val, err := json.Marshal(d) // 此处多加了一层双引号
	if err != nil {
		return val, err
	}
	if len(val) <= 2 {
		return val, errors.New("not valid driver value")
	}
	return val[1 : len(val)-1], nil
}

// Implement sql.Scanner for reading JSONB
func (d *DateTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return d.UnmarshalJSON(v)
	case string:
		return d.UnmarshalJSON([]byte(v))
	case time.Time:
		*d = DateTime(v)
		return nil
	default:
		return errors.New("type assertion to DateTime failed")
	}
}

func (d DateTime) String() string {
	return time.Time(d).Format(dtDefaultFmt)
}
