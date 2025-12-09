package hostlookup

import (
	"context"
	"net"
	"time"
)

func HostLookup(hostname string, timeout time.Duration) bool {
	if timeout == 0 {
		timeout = time.Second
	}
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	_, err := net.DefaultResolver.LookupHost(ctx, hostname)
	return err == nil
}
