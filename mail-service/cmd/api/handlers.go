package main

import (
	"net/http"

	"github.com/tsawler/toolbox"
)

type mailMessage struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {

	tool := toolbox.Tools{}
	var requestPayload mailMessage
	err := tool.ReadJSON(w, r, &requestPayload)
	if err != nil {
		tool.ErrorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		tool.ErrorJSON(w, err)
		return
	}

	payload := toolbox.JSONResponse{
		Error:   false,
		Message: "sent to  " + requestPayload.To,
	}

	tool.WriteJSON(w, http.StatusAccepted, payload)

}
