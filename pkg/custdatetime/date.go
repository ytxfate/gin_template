package custdatetime

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Date time.Time

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	t, err := time.Parse(time.DateOnly, string(data[1:len(data)-1]))
	if err != nil {
		t, err = time.Parse("20060102", string(data[1:len(data)-1]))
		if err != nil {
			return err
		}
	}
	*d = Date(t)
	return
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", time.Time(d).Format(time.DateOnly))), nil
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
	s, ok := value.(string)
	var tm time.Time
	var err error
	if ok {
		tm, err = time.Parse("20060102", s)
		if err == nil {
			*d = Date(tm)
			return nil
		}
	}
	tm, ok = value.(time.Time)
	if !ok {
		return errors.New("type assertion to time.Time failed")
	}
	*d = Date(tm)
	return nil
}

func (d Date) String() string {
	return time.Time(d).Format(time.DateOnly)
}
