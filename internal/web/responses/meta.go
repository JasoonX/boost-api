package responses

import "encoding/json"

// Reason is a struct that contains the reason of a response.
// e.g. "missing authorization header", this should not be used for
// user messages, and in general contains information that is for devs only
// for backend errors that should be processed of front and shown to user please use
// errors.JSONServerErrors
type Reason struct {
	Reason string `json:"reason"`
}

func (r Reason) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Reason string `json:"reason"`
	}{
		Reason: r.Reason,
	})
}
