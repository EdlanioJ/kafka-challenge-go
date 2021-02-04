package model

import "encoding/json"

type Receiver struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (r *Receiver) ToJson() ([]byte, error) {
	out, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *Receiver) ParserJson(data []byte) error {
	err := json.Unmarshal(data, r)
	if err != nil {
		return err
	}

	return nil
}
