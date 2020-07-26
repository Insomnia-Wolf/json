package json

import (
	//std necessities
	"encoding/json"
	"net/http"
)

func Marshal(source interface{}, pretty bool) ([]byte, error) {
	if pretty {
		return json.MarshalIndent(source, "", "\t")
	}
	return json.Marshal(source)
}

func Unmarshal(dataJSON []byte, target interface{}) (err error) {
	err = json.Unmarshal(dataJSON, target)
	return
}

func UnmarshalBody(body *http.Response, target interface{}) (err error) {
	defer body.Body.Close()
	return json.NewDecoder(body.Body).Decode(target)
}