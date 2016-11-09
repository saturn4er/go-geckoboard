package geckoboard

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Application struct {
	ApiToken string
}
type Error struct {
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}
func (ga *Application) request(method, endooint string, data interface{}) (err error) {
	var requestBody []byte
	if data != nil {
		requestBody, err = json.Marshal(data)
		if err != nil {
			return
		}
	}
	request, err := http.NewRequest(method, "https://api.geckoboard.com"+endooint, bytes.NewReader(requestBody))
	if err != nil {
		return
	}
	request.Header.Add("Content-Type", "application/json")
	request.SetBasicAuth(ga.ApiToken, "")
	client := &http.Client{}
	resp, err := client.Do(request)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		errResp := struct {
			Error Error
		}{}
		err = json.Unmarshal(body, &errResp)
		if err != nil {
			return
		}
		err = errResp.Error
		return
	}
	return nil
}

type Dataset struct {
	Fields   Fields   `json:"fields"`
	UniqueBy []string `json:"unique_by"`
}

func (ga *Application) CreateDataset(id string, dataset Dataset) error {
	if id == "" {
		return errors.New("bad id")
	}
	return ga.request("PUT", "/datasets/"+id, dataset)
}

type Data struct {
	Data []map[string]interface{} `json:"data"`
}

func (ga *Application) UpdateDatasetData(id string, data []map[string]interface{}) error {
	return ga.request("PUT", "/datasets/"+id+"/data", Data{Data: data})
}
func (ga *Application) AddDataToDataset(id string, data []map[string]interface{}) error {
	return ga.request("POST", "/datasets/"+id+"/data", Data{Data: data})
}
func (ga *Application) DeleteDataset(id string) error {
	return ga.request("DELETE", "/datasets/"+id, nil)
}
