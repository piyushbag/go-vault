package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Request struct {
	Action string     `json:"action"`
	Auth   Auth       `json:"auth,omitempty"`
	Log    LogPayload `json:"log,omitempty"`
}

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (c *Config) Broker(w http.ResponseWriter, r *http.Request) {
	var res jsonResponse
	res.Error = false
	res.Message = "Broker service is running"
	res.Data = "Service 1 is running"

	_ = c.writeJSON(w, http.StatusOK, res)
}

// communicate with the authentication service (internal service)
func (c *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var req Request
	// 1. read the request
	err := c.readJSON(w, r, &req)
	if err != nil {
		c.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// 2. check the action
	switch req.Action {
	case "auth":
		c.Authenticate(w, req.Auth)
	case "log":
		c.Log(w, req.Log)
	default:
		c.errorJSON(w, errors.New("unknown action"))
	}
}

func (c *Config) Authenticate(w http.ResponseWriter, auth Auth) {
	// 1. Marshal the json we will send to the authentication service by using the Auth payload
	data, _ := json.MarshalIndent(auth, "", "\t")

	// 2. convert the JSON payload to a request to the authentication service
	request, err := http.NewRequest("POST", "http://authentication/authenticate", bytes.NewBuffer(data))
	if err != nil {
		c.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// 3. create a new http client to send the request
	client := &http.Client{}
	// 4. send the request to the authentication service
	response, err := client.Do(request)
	if err != nil {
		c.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// 5. check the response from the authentication service
	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		c.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		c.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	var authResponse jsonResponse
	// 6. decode the response from the authentication service
	err = json.NewDecoder(response.Body).Decode(&authResponse)
	if err != nil {
		c.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// 7. check if the authentication was successful
	if authResponse.Error {
		c.errorJSON(w, errors.New(authResponse.Message), http.StatusUnauthorized)
		return
	}

	// 8. send the response back to the client
	var res jsonResponse
	res.Error = false
	res.Message = "Authenticated!!!"
	res.Data = authResponse.Data

	c.writeJSON(w, http.StatusOK, res)
}

func (c *Config) Log(w http.ResponseWriter, entry LogPayload) {
	// 1. marshal the log payload
	data, _ := json.MarshalIndent(entry, "", "\t")

	// 2. convert the json payload to a request to the authentication service
	logURL := "http://logger/log"
	request, err := http.NewRequest("POST", logURL, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("test...")
		c.errorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	// 3. create a new http client to send the request
	client := &http.Client{}
	// 4. send the request to the authentication service
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("test2...")
		c.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// 5. check the response from the authentication service
	// make sure we get back the correct status code
	fmt.Println(response.StatusCode, "code..")
	if response.StatusCode != http.StatusAccepted {
		fmt.Println("test3....")
		c.errorJSON(w, err)
		return
	}

	// BECAUSE ITS COMMUNICATING WITH ANOTHER SERVICE
	// AND NOT DIRECTLY COMMUNICATING TO FRONT END HENCE NOT REQUIRED TO
	// BUILD JSON RESPONSE

	// var logResponse jsonResponse
	// // 6. decode the response from the authentication service
	// err = json.NewDecoder(response.Body).Decode(&logResponse)
	// if err != nil {
	// 	c.errorJSON(w, err, http.StatusInternalServerError)
	// 	return
	// }

	// // 7. check if the authentication was successful
	// if logResponse.Error {
	// 	c.errorJSON(w, errors.New(logResponse.Message), http.StatusUnauthorized)
	// 	return
	// }

	// 8. write back response
	var resp jsonResponse
	resp.Error = false
	resp.Message = "connected to logger service"

	c.writeJSON(w, http.StatusAccepted, resp)
}
