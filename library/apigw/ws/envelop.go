package ws

import "encoding/json"

type InputEnvelop struct {
	Echo bool            `json:"echo"`
	Type int             `json:"type"`
	Data json.RawMessage `json:"data"`
}

func (e *InputEnvelop) Decode(data []byte) (*InputEnvelop, error) {
	err := json.Unmarshal(data, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

type OutputEnvelop struct {
	Type     int             `json:"type"`
	Data     json.RawMessage `json:"data"`
	Received int64           `json:"received"`
}

func (e *OutputEnvelop) Encode() ([]byte, error) {
	return json.Marshal(e)
}
