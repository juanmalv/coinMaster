package utils

import (

	"io/ioutil"

	"net/http"
)

func ReadBody(resp *http.Response ) ([]byte, error){
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		return nil,err
	}
	return bodyBytes, nil
}
