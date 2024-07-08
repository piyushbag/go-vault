package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var req request
	err := c.readJSON(w, r, &req)
	if err != nil {
		c.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user credentials
	user, err := c.Models.User.GetByEmail(req.Email)
	if err != nil {
		c.errorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	// check the password
	valid, err := user.PasswordMatches(req.Password)
	if err != nil || !valid {
		c.errorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	// log authentication
	err = c.logRequest("authentication", fmt.Sprintf("%s logged in", user.Email))
	if err != nil {
		c.errorJSON(w, err)
		return
	}

	var resp response
	resp.Error = false
	resp.Message = "authenticated"
	resp.Data = user

	c.writeJSON(w, http.StatusAccepted, resp)
}

func (c *Config) logRequest(name, data string) error {
	var entry struct {
		Name string
		Data string
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	logServiceURL := "http://logger/log"
	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil

}
