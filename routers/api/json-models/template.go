package json_models

type JSONData interface {

}

// generic response
// swagger:response statusResponse
type RespStatus struct {
	// in:body
	Body JSONTemplate
}

type JSONTemplate struct {
	// response code
	Code int `json:"code"`
	// response message contains status message
	Msg string `json:"message"`
	// contains data, for status response, no data contains
	JSONData JSONData `json:"data"`
}
