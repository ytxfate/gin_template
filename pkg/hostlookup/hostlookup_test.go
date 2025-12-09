package hostlookup

import (
	"fmt"
	"testing"
)

func TestHostLookup(t *testing.T) {
	stat := HostLookup("localhost", 0)
	fmt.Println(stat)
}
