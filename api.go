package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/jforcode/Go-DeepError"
)

type IApi interface {
	CreateEvent(event *Event) (string, error)
	GetAllEvents() ([]*Event, error)
	GetEvent(eventID string) (*Event, error)
}

type Api struct {
	url    string
	client *http.Client
}

const (
	createEP = "/event"
	getAllEP = "/events"
	getEP    = "/events/:eventID"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err ApiError) Error() string {
	return strconv.Itoa(err.Code) + ":" + err.Message
}

type ApiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   ApiError    `json:"error"`
}

// EventIDResponse represents the response to send back to client, in case of create event
type EventIDResponse struct {
	EventID string `json:"eventID"`
}

// EventResponse represents the response to send to client, in case of a get event
type EventResponse struct {
	Event *Event `json:"event"`
}

// EventsResponse represents the response to send to client, in case of a get all events call
type EventsResponse struct {
	Events []*Event `json:"events"`
}

// CreateEvent uses the API to create an event
func (api *Api) CreateEvent(event *Event) (string, error) {
	fn := "CreateEvent"

	body, err := json.Marshal(event)
	if err != nil {
		return "", deepError.New(fn, "marshal", err)
	}

	url := api.url + createEP
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", deepError.New(fn, "new request", err)
	}

	readInto := &EventIDResponse{}
	err = api.makeRequestAndReadJSON(req, readInto)
	if err != nil {
		return "", deepError.New(fn, "requesting and reading", err)
	}

	return readInto.EventID, nil
}

// GetAllEvents uses the API to get all events
func (api *Api) GetAllEvents() ([]*Event, error) {
	fn := "GetAllEvents"

	url := api.url + getAllEP
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, deepError.New(fn, "new request", err)
	}

	readInto := &EventsResponse{}
	err = api.makeRequestAndReadJSON(req, readInto)
	if err != nil {
		return nil, deepError.New(fn, "requesting and reading", err)
	}

	return readInto.Events, nil
}

// GetEvent uses the API to get the event with the given event ID
func (api *Api) GetEvent(eventID string) (*Event, error) {
	fn := "GetEvent"

	url := strings.Replace(api.url+getEP, ":eventID", eventID, -1)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, deepError.New(fn, "new request", err)
	}

	readInto := &EventResponse{}
	err = api.makeRequestAndReadJSON(req, readInto)
	if err != nil {
		return nil, deepError.New(fn, "requesting and reading", err)
	}

	return readInto.Event, nil
}

func (api *Api) makeRequestAndReadJSON(req *http.Request, readInto interface{}) error {
	fn := "makeRequestAndReadJSON"

	resp, err := api.client.Do(req)
	if err != nil {
		return deepError.New(fn, "making request", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return deepError.New(fn, "Checking status", errors.New("Error: "+resp.Status))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return deepError.New(fn, "read all from response", err)
	}

	apiResp := ApiResponse{}
	err = json.Unmarshal(respBody, &apiResp)
	if err != nil {
		return deepError.New(fn, "unmarshal", err)
	}

	if !apiResp.Success {
		return deepError.New(fn, "unsuccessful request", apiResp.Error)
	}

	dataPart, err := json.Marshal(apiResp.Data)
	if err != nil {
		return deepError.New(fn, "marshalling data part", err)
	}

	err = json.Unmarshal(dataPart, readInto)
	if err != nil {
		return deepError.New(fn, "reading data part into variable", err)
	}

	return nil
}
