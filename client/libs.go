package client

import (
	"github.com/segmentio/objconv/json"
	cerrors "github.com/senfix/fxcm/errors"
	"github.com/senfix/fxcm/response"
)

func validateResponse(data []byte) error {

	header := response.Header{}
	err := json.Unmarshal(data, &header)
	if err != nil {
		return cerrors.UnmarshalError(err)
	}

	if len(header.Response.Error) > 0 {
		return cerrors.ApiError(header.Response.Error)
	}

	if header.Response.Executed == false {
		return cerrors.ExecuteError()
	}

	return nil
}
