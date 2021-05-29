package hpool

import (
	"encoding/json"
)

type responseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type resultStatus struct {
	Status string `json:"status"`
	Result map[string]json.RawMessage
}

type listData struct {
	Total int               `json:"total"`
	List  json.RawMessage `json:"list"`
}

type jsonResponse struct {
	Code int        `json:"code"`
	Message string `json:"message"`
	Data json.RawMessage `json:"data"`

}

//func (rs *resultStatus) UnmarshalJSON_(b []byte) error {
//	var r map[string]json.RawMessage
//	if err := json.Unmarshal(b, &r); err != nil {
//		return err
//	}
//	s, ok := r["status"]
//	if ok {
//		_ = json.Unmarshal(s, rs.Status)
//		delete(r, "status")
//	}
//	rs.Result = r
//	return nil
//}
