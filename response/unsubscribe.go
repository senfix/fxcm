package response

import "github.com/senfix/fxcm"

type Unsubscribe struct {
	Header
	fxcm.Unsubscribe
}
