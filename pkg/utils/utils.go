package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Parsing body failed")
		return
	}

	err = json.Unmarshal([]byte(body), x)

	if err != nil {
		fmt.Println("Failed to unmarshal body")
		return
	}
}
