package dans

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type QueryListJob struct {
	Page        string `json:"page"`
	Description string `json:"description"`
	FullTime    string `json:"full_time"`
}

type QueryJobDetail struct {
	Id string `json:"id"`
}

func GetListJob(qParams QueryListJob) (*[]DansModel, error) {
	url := "http://dev3.dansmultipro.co.id/api/recruitment/positions.json"
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	q := req.URL.Query()
	q.Add(`page`, qParams.Page)
	q.Add(`description`, qParams.Description)
	q.Add(`full_time`, qParams.FullTime)
	req.URL.RawQuery = q.Encode()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return nil, err
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed request")
	}
	defer resp.Body.Close()
	result := new([]DansModel)
	if err := json.Unmarshal(resBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetListJobDetail(qParams string) (*DansModel, error) {
	url := "http://dev3.dansmultipro.co.id/api/recruitment/positions/"
	req, _ := http.NewRequest(http.MethodGet, url + qParams, nil)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return nil, err
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed request")
	}
	defer resp.Body.Close()
	result := new(DansModel)
	if err := json.Unmarshal(resBody, result); err != nil {
		return nil, err
	}
	return result, nil
}
