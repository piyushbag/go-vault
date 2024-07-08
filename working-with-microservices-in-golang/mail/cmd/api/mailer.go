package main

import "bytes"

type Mail struct {
	Domain      string `json:"domain"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Encryption  string `json:"encryption"`
	FromAddress string `json:"from_address"`
	FromName    string `json:"from_name"`
}

type Message struct {
	From        string                 `json:"from"`
	FromName    string                 `json:"from_name"`
	To          string                 `json:"to"`
	Subject     string                 `json:"subject"`
	Attachments []string               `json:"attachments"`
	Data        interface{}            `json:"data"`
	DataMap     map[string]interface{} `json:"data_map"`
}

func (m *Mail) Send(msg Message) error {
	if msg.From == "" {
		msg.From = m.FromAddress
	}

	if msg.FromName == "" {
		msg.FromName = m.FromName
	}

	data := map[string]interface{}{
		"message": msg.Data,
	}

	msg.DataMap = data
	formattedMsg, err := m.buildHTMLMessage(msg)
}

func (m *Mail) buildHTMLMessage(msg Message) (string, error) {
	var body bytes.Buffer
	body.WriteString("<html><body>")
	body.WriteString("<h1>" + msg.Subject + "</h1>")
	body.WriteString("<p>From: " + msg.FromName + " <" + msg.From + "></p>")
	body.WriteString("<p>To: " + msg.To + "</p>")
	body.WriteString("<hr>")
	body.WriteString("<p>" + msg.Data.(string) + "</p>")
	body.WriteString("</body></html>")

	return body.String(), nil
}
