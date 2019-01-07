package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{}
	api := &Api{
		url:    "http://localhost:4000",
		client: client,
	}

	userData, err := ParseCmd()
	if err != nil {
		panic(err)
	}

	if userData.Command == cmdList {
		handleList(api, userData.ListData)
	} else if userData.Command == cmdCreate {
		handleCreate(api, userData.CreateData)
	}
}

func handleList(api *Api, listData *ListData) {
	if listData.EventID == "" {
		events, err := api.GetAllEvents()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, event := range events {
			fmt.Println(event.ToJSON())
		}

	} else {
		event, err := api.GetEvent(listData.EventID)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(event.ToJSON())
	}
}

func handleCreate(api *Api, createData *CreateData) {
	eventTags := make([]*EventTag, len(createData.Tags))
	for _, cmdTag := range createData.Tags {
		eventTags = append(eventTags, &EventTag{Value: cmdTag})
	}

	eventType := &EventType{Value: "single"}

	event := &Event{
		Title:         createData.Title,
		Note:          createData.Desc,
		Tags:          eventTags,
		Type:          eventType,
		UserCreatedAt: time.Now().UTC(),
	}

	eventID, err := api.CreateEvent(event)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Created event with ID: " + eventID)
}
