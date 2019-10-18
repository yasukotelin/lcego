package main

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type License struct {
	Key string `json:"key"`
	Name string `json:"name"`
	SpdxId string `json:"spdx_id"`
	Url string `json:"url"`
	NodeId string `json:"node_id"`
}

type LicenseDetail struct {
	Key string `json:"key"`
	Name string `json:"name"`
	SpdxId string `json:"spdx_id"`
	Url string `json:"url"`
	NodeId string `json:"node_id"`
	HtmlUrl string `json:"html_url"`
	Description string `json:"description"`
	Implementation string `json:"implementation"`
	Permissions []string `json:"permissions"`
	Conditions []string `json:"conditions"`
	Limitations []string `json:"limitations"`
	Body string `json:"body"`
	Featured bool `json:"featured"`
}

func getLicenses() ([]License, error) {
	url := "https://api.github.com/licenses"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var body []License
	if err := json.Unmarshal(bytes, &body); err != nil {
		return nil, err
	}

	return body, nil
}

func getLicenseDetail(license *License) (*LicenseDetail, error) {
	res, err := http.Get(license.Url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	body := new(LicenseDetail)
	if err := json.Unmarshal(bytes, body); err != nil {
		return nil, err
	}

	return body, nil
}
