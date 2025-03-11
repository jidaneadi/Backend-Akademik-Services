package helpers

import "backend-sia/models/response"

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
	// return nil
}

func ErrorBadRequest(err error) interface{} {
	if err != nil {
		return response.Error{
			Meta: response.Meta{
				Code:    400,
				Status:  false,
				Message: err.Error(),
			},
		}
	}
	return nil
}
