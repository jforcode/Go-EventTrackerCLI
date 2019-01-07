package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/jforcode/Go-DeepError"
)

const (
	tagSeparator = "#"
	cmdList      = "list"
	cmdCreate    = "create"
)

// TagFlags is the custom interface to get tags from command line
type TagFlags []string

// String is part of the Value interface of flags.
// used to format the flag's value
func (tagFlags *TagFlags) String() string {
	if tagFlags != nil {
		return fmt.Sprint(*tagFlags)
	}

	return ""
}

// Set is part of the value interface of flags
// used to set the value from cmd to the interface in code
func (tagFlags *TagFlags) Set(value string) error {
	if tagFlags != nil {
		for _, val := range strings.Split(value, tagSeparator) {
			*tagFlags = append(*tagFlags, val)
		}
	}

	return nil
}

// UserData represents the entire command line arguments
type UserData struct {
	Command    string
	ListData   *ListData
	CreateData *CreateData
}

// ToJSON prints the use data in JSON
func (userData *UserData) ToJSON() string {
	fn := "ListData.ToString"

	jsonBytes, err := json.MarshalIndent(userData, "", "    ")
	if err != nil {
		return deepError.New(fn, "marshalling", err).Error()
	}

	return string(jsonBytes)
}

// ListData represents the data if command is list
type ListData struct {
	EventID string
}

// ToJSON prints the list data in JSON
func (listData *ListData) ToJSON() string {
	fn := "ListData.ToString"

	jsonBytes, err := json.MarshalIndent(listData, "", "    ")
	if err != nil {
		return deepError.New(fn, "marshalling", err).Error()
	}

	return string(jsonBytes)
}

// CreateData represents the data if command is create
type CreateData struct {
	Title string
	Desc  string
	Tags  []string
}

// ToJSON prints the create data in JSON
func (createData *CreateData) ToJSON() string {
	fn := "CreateData.ToString"

	jsonBytes, err := json.MarshalIndent(createData, "", "    ")
	if err != nil {
		return deepError.New(fn, "marshalling", err).Error()
	}

	return string(jsonBytes)
}

var eventIDFlag string
var titleFlag string
var descFlag string
var tagFlags *TagFlags

var listFlag bool
var createFlag bool

// ParseCmd parses user entered command line data
func ParseCmd() (*UserData, error) {
	fn := "ParseCmd"
	flag.StringVar(&eventIDFlag, "eventID", "", "EventID to fetch the event for")
	flag.StringVar(&eventIDFlag, "id", "", "EventID to fetch the event for (shorthand)")

	var tagFlags *TagFlags
	flag.StringVar(&titleFlag, "title", "", "The title of the event")
	flag.StringVar(&descFlag, "desc", "", "The actual event description")
	flag.Var(tagFlags, "tags", "The tags for the event")

	flag.BoolVar(&listFlag, "list", false, "Use this flag to list all events")
	flag.BoolVar(&listFlag, "l", false, "Use this flag to list all events (shorthand)")
	flag.BoolVar(&createFlag, "create", false, "Use this flag to create an event")
	flag.BoolVar(&createFlag, "c", false, "Use this flag to create an event (shorthand)")

	flag.Parse()
	args := flag.Args()

	userData := &UserData{}

	if listFlag {
		listData, err := ParseListData(args, []interface{}{eventIDFlag})
		if err != nil {
			return nil, deepError.New(fn, "parsing list data", err)
		}

		userData.Command = cmdList
		userData.ListData = listData

		return userData, nil
	}

	if createFlag {
		createData, err := ParseCreateData(args, []interface{}{titleFlag, descFlag, tagFlags})
		if err != nil {
			return nil, deepError.New(fn, "parsing create data", err)
		}

		userData.Command = cmdCreate
		userData.CreateData = createData

		return userData, nil
	}

	return nil, deepError.New(fn, "Invalid command", nil)
}

// ParseListData parses the data if command is list
func ParseListData(args []string, flags []interface{}) (*ListData, error) {
	eventIDFlag := flags[0].(string)
	listData := &ListData{}

	listData.EventID = eventIDFlag
	if listData.EventID == "" {
		if len(args) > 0 {
			listData.EventID = args[0]
			args = args[1:]
		}
	}

	return listData, nil
}

// ParseCreateData parses the data if command is create.
func ParseCreateData(args []string, flags []interface{}) (*CreateData, error) {
	fn := "ParseCreateData"

	titleFlag := flags[0].(string)
	descFlag := flags[1].(string)
	tagFlags := flags[2].(*TagFlags)

	createData := &CreateData{}

	createData.Title = titleFlag
	if createData.Title == "" {
		if len(args) > 0 {
			createData.Title = args[0]
			args = args[1:]
		} else {
			return nil, deepError.New(fn, "No title available either as a FLAG or as an ARG", nil)
		}
	}

	createData.Desc = descFlag
	if createData.Desc == "" {
		if len(args) > 0 {
			createData.Desc = args[0]
			args = args[1:]
		}
	}

	createData.Tags = make([]string, 0)
	if tagFlags != nil {
		createData.Tags = []string(*tagFlags)
	}
	for _, arg := range args {
		createData.Tags = append(createData.Tags, strings.Split(arg, tagSeparator)...)
	}

	return createData, nil
}
