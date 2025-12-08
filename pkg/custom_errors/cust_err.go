package customerrors

import commresp "gin_template/pkg/comm_resp"

type CustErr struct {
	Code commresp.StatusCode
	Err  error
}

func (msg CustErr) Error() string {
	return msg.Err.Error()
}
