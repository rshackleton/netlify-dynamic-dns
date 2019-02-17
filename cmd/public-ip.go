package main

import (
	"encoding/json"
)

const ipv4ApiEndpoint = "https://api.ipify.org/?format=json"

type apiResponse struct {
	IP string
}

func getPublicIP() (ipv4 string, err error) {
	ipv4res, err := pingAPI(ipv4ApiEndpoint)
	if err != nil {
		return "", err
	}
	return ipv4res.IP, nil
}

func pingAPI(url string) (apiResponse, error) {
	res, err := newHTTPClient().Get(url)
	if err != nil {
		return apiResponse{}, err
	}
	var response apiResponse
	return response, json.NewDecoder(res.Body).Decode(&response)
}
