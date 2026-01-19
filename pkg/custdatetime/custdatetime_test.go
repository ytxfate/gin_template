package custdatetime

import (
	"fmt"
	"testing"
	"time"
)

func TestDateTime3MilliScan(t *testing.T) {
	var d DateTime3Milli
	err := d.Scan("20251023142536789")
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	err = d.Scan("\"20251023142536786\"")
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	err = d.Scan([]byte("\"20251023142536786\""))
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	err = d.Scan(time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

	// 反例
	err = d.Scan(Date(time.Now()))
	if err == nil {
		panic("Date not support")
	}

	err = d.Scan("20252023142536789")
	if err == nil {
		panic("20252023142536789 not Time")
	}

}

func TestDateTime3MilliUnmarshalJSON(t *testing.T) {
	var d DateTime3Milli
	err := d.UnmarshalJSON([]byte("20251023142536789"))
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	err = d.UnmarshalJSON([]byte("\"20251023142536786\""))
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
}

func TestDateTime3MilliMarshalJSON(t *testing.T) {
	var d DateTime3Milli = DateTime3Milli(time.Now())
	s, err := d.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}

func TestDateTime3MilliValue(t *testing.T) {
	var d DateTime3Milli = DateTime3Milli(time.Now())
	s, err := d.Value()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", s)
}

func TestDateTimeScan(t *testing.T) {
	var d DateTime
	err := d.Scan("20251023142536")
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	err = d.Scan("\"20251023142535\"")
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	err = d.Scan([]byte("\"20251023142536\""))
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	err = d.Scan(time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

	// 反例
	err = d.Scan(Date(time.Now()))
	if err == nil {
		panic("Date not support")
	}

	err = d.Scan("20252023142536")
	if err == nil {
		panic("20252023142536 not Time")
	}

}

func TestDateTimeUnmarshalJSON(t *testing.T) {
	var d DateTime
	err := d.UnmarshalJSON([]byte("20251023142536"))
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	err = d.UnmarshalJSON([]byte("\"20251023142536\""))
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
}

func TestDateTimeMarshalJSON(t *testing.T) {
	var d DateTime = DateTime(time.Now())
	s, err := d.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}

func TestDateTimeValue(t *testing.T) {
	var d DateTime = DateTime(time.Now())
	s, err := d.Value()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", s)
}

func TestDateScan(t *testing.T) {
	var d Date
	err := d.Scan("20251023")
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	err = d.Scan("\"20251024\"")
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	err = d.Scan([]byte("\"20251026\""))
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	err = d.Scan(time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

	// 反例
	err = d.Scan(DateTime(time.Now()))
	if err == nil {
		panic("Date not support")
	}

	err = d.Scan("20252023")
	if err == nil {
		panic("20252023 not Time")
	}

}

func TestDateUnmarshalJSON(t *testing.T) {
	var d Date
	err := d.UnmarshalJSON([]byte("20251023"))
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	err = d.UnmarshalJSON([]byte("\"20251024\""))
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
}

func TestDateMarshalJSON(t *testing.T) {
	var d Date = Date(time.Now())
	s, err := d.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}

func TestDateValue(t *testing.T) {
	var d Date = Date(time.Now())
	s, err := d.Value()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", s)
}
