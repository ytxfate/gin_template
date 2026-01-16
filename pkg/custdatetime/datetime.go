package custdatetime

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type DateTime time.Time

func (d *DateTime) UnmarshalJSON(data []byte) (err error) {
	t, err := time.Parse(time.DateTime, string(data[1:len(data)-1]))
	if err != nil {
		t, err = time.Parse("20060102150405", string(data[1:len(data)-1]))
		if err != nil {
			return err
		}
	}
	*d = DateTime(t)
	return
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", time.Time(d).Format(time.DateTime))), nil
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
	tm, ok := value.(time.Time)
	if !ok {
		return errors.New("type assertion to time.Time failed")
	}
	*d = DateTime(tm)
	return nil
}

func (d DateTime) String() string {
	return time.Time(d).Format(time.DateTime)
}
