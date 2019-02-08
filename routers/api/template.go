package api

type JSONData interface {

}

type JSONTemplate struct {
	Code int `json:"code"`
	Msg string `json:"message"`
	JSONData JSONData `json:"data"`
}

//func (f *JSONTemplate) MarshalJSON() ([]byte, error) {
//	type tmp *JSONTemplate
//	g := tmp(f)
//
//	b1, err := json.Marshal(g)
//	if err != nil {
//		return nil, err
//	}
//
//	b2, err := json.Marshal(g.JSONData)
//	if err != nil {
//		return nil, err
//	}
//
//	s1 := string(b1[:len(b1)-1])
//	s2 := string(b2[1:])
//
//	return []byte(s1 + ", " + s2), nil
//}