package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Query struct {
	From        uint64                   `json:"from,omitempty"`
	Size        uint64                   `json:"size,omitempty"`
	QuerySearch interface{}              `json:"query,omitempty"`
	Sort        []map[string]interface{} `json:"sort,omitempty"`
}

// PutData -
func PutData(address string, data interface{}) error {
	client := http.Client{}
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("POST", address, bytes.NewBuffer(body))
	req.Header.Set("Content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return fmt.Errorf("Failed to put data on elastic search")
	}

	return nil
}

// GetData -
func GetData(address string, query interface{}) ([]byte, error) {
	client := http.Client{}
	var data []byte
	body, err := json.Marshal(query)
	if err != nil {
		return data, err
	}
	req, _ := http.NewRequest("POST", address, bytes.NewBuffer(body))
	req.Header.Set("Content-type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return data, err
	}
	data, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println(res.StatusCode, string(data))
		return data, fmt.Errorf("Failed to get data on elastic search")
	}

	return data, err
}
