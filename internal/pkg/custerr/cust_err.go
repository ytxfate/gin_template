package custerr

import "gin_template/internal/pkg/commresp"

type CustErr struct {
	Code commresp.StatusCode
	Err  error
}

func (msg CustErr) Error() string {
	return msg.Err.Error()
}
