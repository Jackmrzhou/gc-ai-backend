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
	// Example: 400
	Code int `json:"code"`
	// response message
	// Example: invalid parameters
	Msg string `json:"message"`
	// Example: nil
	JSONData JSONData `json:"data"`
}
