package auth

import (
	"errors"
	"net/http"
	"strings"
)

//GetAPIKey extracts an API key from the header
//Example :
//Authorization : APIKey : {insert your apikey here}
func GetAPIKey(headers http.Header) (string,error){
	vals := headers.Get("Authorization")

	if vals == "" {
		return "",errors.New("no authentication info found")
	}

	val:= strings.Split(vals," ")

	if len(val) !=2 {
		return "",errors.New("wrong format of the APIKey struct")
	}

	if val[0] != "APIKey" {
		return "",errors.New("wrong format of APIKey")	
	} 

	return val[1],nil
}