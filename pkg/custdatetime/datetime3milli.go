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
	dt3mDefaultFmt = "2006-01-02 15:04:05.000"
	dt3mCompactFmt = "20060102150405.000"
)

type DateTime3Milli time.Time

func (d *DateTime3Milli) UnmarshalJSON(data []byte) (err error) {
	ds := strings.Trim(string(data), "\"")
	t, err := time.Parse(dt3mDefaultFmt, ds)
	if err != nil {
		// 移除非必要字符
		ds = strings.ReplaceAll(ds, "-", "")
		ds = strings.ReplaceAll(ds, ":", "")
		ds = strings.ReplaceAll(ds, " ", "")
		ds = strings.ReplaceAll(ds, "T", "")
		ds = strings.ReplaceAll(ds, ".", "")
		ds += "00000000000000000"
		ds = ds[:14] + "." + ds[14:17] // 日期处理成 20060102150405.000 格式
		t, err = time.Parse(dt3mCompactFmt, ds)
		if err != nil {
			return err
		}
	}
	*d = DateTime3Milli(t)
	return
}

func (d DateTime3Milli) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", time.Time(d).Format(dt3mDefaultFmt))), nil
}

// Implement driver.Valuer for saving JSONB
func (d DateTime3Milli) Value() (driver.Value, error) {
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
func (d *DateTime3Milli) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return d.UnmarshalJSON(v)
	case string:
		return d.UnmarshalJSON([]byte(v))
	case time.Time:
		*d = DateTime3Milli(v)
		return nil
	default:
		return errors.New("type assertion to DateTime3Milli failed")
	}
}

func (d DateTime3Milli) String() string {
	return time.Time(d).Format(dt3mDefaultFmt)
}
