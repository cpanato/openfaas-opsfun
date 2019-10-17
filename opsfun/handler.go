package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

func Handle(req handler.Request) (handler.Response, error) {
	var err error

	var query *url.Values
	if len(req.Body) > 0 {
		q, err := url.ParseQuery(string(req.Body))
		if err != nil {
			return handler.Response{}, err
		}
		query = &q
	}

	command := query.Get("command")
	message := processCommand(command)

	return handler.Response{
		Body:       []byte(message),
		StatusCode: http.StatusOK,
		Header: map[string][]string{
			"Content-Type": []string{"application/json"},
		},
	}, err
}

func processCommand(command string) string {
	if len(command) > 0 {

		switch command {
		case "/magic":
			gif := "in response to `/magic`\n\nhttps://thumbs.gfycat.com/AmbitiousUnselfishHake-size_restricted.gif"
			return generateStandardSlashResponse(gif, "in_channel", "Magic")
		case "/this-is-fine":
			gif := "in response to `/this-is-fine`\n\nhttps://thumbs.gfycat.com/AngryGoldenIsabellinewheatear-size_restricted.gif"
			return generateStandardSlashResponse(gif, "in_channel", "This is Fine")
		case "/everything-is-on-fire":
			gif := "in response to `/everything-is-on-fire`\n\nhttps://media.giphy.com/media/13HgwGsXF0aiGY/giphy.gif"
			return generateStandardSlashResponse(gif, "in_channel", "Fire!")
		}
	}

	return ""
}

type MMSlashResponse struct {
	ResponseType string `json:"response_type,omitempty"`
	Username     string `json:"username,omitempty"`
	IconUrl      string `json:"icon_url,omitempty"`
	Channel      string `json:"channel,omitempty"`
	Text         string `json:"text,omitempty"`
	GotoLocation string `json:"goto_location,omitempty"`
}

func generateStandardSlashResponse(text, respType, username string) string {
	response := MMSlashResponse{
		ResponseType: respType,
		Text:         text,
		GotoLocation: "",
		Username:     username,
	}

	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Unable to marshal response")
		return ""
	}
	return string(b)
}
