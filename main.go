package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
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

		for i := len(events) - 1; i >= 0; i-- {
			event := events[i]
			fmt.Println(getEventMini(event))
		}

	} else {
		event, err := api.GetEvent(listData.EventID)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(getEventMini(event))
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

var timeColor = color.New(color.FgGreen)
var titleColor = color.New(color.FgYellow)
var tagColor = color.New(color.FgCyan)
var idColor = color.New(color.FgGreen)

func getEventMini(event *Event) string {
	if event == nil {
		return "<nil>"
	}

	var miniStrBldr strings.Builder

	timePart := timeColor.Sprint(event.UserCreatedAt.Format(time.RFC1123Z))
	titlePart := titleColor.Sprint(event.Title)
	miniStrBldr.WriteString(fmt.Sprintf("%s %s ", timePart, titlePart))

	for _, tag := range event.Tags {
		if tag != nil {
			tagPart := tagColor.Sprint(tagSeparator + tag.Value)
			miniStrBldr.WriteString(tagPart + " ")
		}
	}

	eventIDPart := idColor.Sprint(event.ID)
	miniStrBldr.WriteString(fmt.Sprintf("(%s)", eventIDPart))

	return miniStrBldr.String()
}

func getEventFull(event *Event) string {
	if event == nil {
		return "<nil>"
	}

	return event.ToJSON()
}
