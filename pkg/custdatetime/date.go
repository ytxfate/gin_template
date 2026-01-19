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
	dDefaultFmt = time.DateOnly
	dCompactFmt = "20060102"
)

type Date time.Time

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	ds := strings.Trim(string(data), "\"")
	t, err := time.Parse(dDefaultFmt, ds)
	if err != nil {
		t, err = time.Parse(dCompactFmt, ds)
		if err != nil {
			return err
		}
	}
	*d = Date(t)
	return
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", time.Time(d).Format(dDefaultFmt))), nil
}

// Implement driver.Valuer for saving JSONB
func (d Date) Value() (driver.Value, error) {
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
func (d *Date) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return d.UnmarshalJSON(v)
	case string:
		return d.UnmarshalJSON([]byte(v))
	case time.Time:
		*d = Date(v)
		return nil
	default:
		return errors.New("type assertion to Date failed")
	}
}

func (d Date) String() string {
	return time.Time(d).Format(dDefaultFmt)
}
