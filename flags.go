package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/jforcode/Go-DeepError"
)

const tagSeparator = ";"

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

// UserData is the model we use to represent the user commnad line input, instead of using separate variables
type UserData struct {
	Title string   `json:"title"`
	Desc  string   `json:"desc"`
	Tags  []string `json:"tags"`
}

// ToJSON returns a json string of user data
func (userData *UserData) ToJSON() string {
	fn := "UserData.ToString"

	fmt.Println(userData)
	jsonBytes, err := json.MarshalIndent(userData, "", "    ")
	if err != nil {
		return deepError.New(fn, "marshalling", err).Error()
	}

	return string(jsonBytes)
}

// ParseCmd parses user entered command line data
func ParseCmd() (*UserData, error) {
	fn := "ParseCmd"

	var tagFlags *TagFlags
	titleFlag := flag.String("title", "", "The title of the event")
	descFlag := flag.String("desc", "", "The actual event description")
	flag.Var(tagFlags, "tags", "The tags for the event")

	flag.Parse()
	args := flag.Args()
	lenArgs := len(args)

	userData := &UserData{}

	userData.Title = *titleFlag
	if strings.Compare(userData.Title, "") == 0 {
		if lenArgs > 0 {
			userData.Title = args[0]
			args = args[1:]
		} else {
			return nil, deepError.New(fn, "No title available either as a FLAG or as an ARG", nil)
		}
	}

	userData.Desc = *descFlag
	if strings.Compare(userData.Desc, "") == 0 {
		if lenArgs > 0 {
			userData.Desc = args[0]
			args = args[1:]
		}
	}

	userData.Tags = make([]string, 0)
	if tagFlags != nil {
		userData.Tags = []string(*tagFlags)
	}
	for _, arg := range args {
		userData.Tags = append(userData.Tags, strings.Split(arg, tagSeparator)...)
	}

	return userData, nil
}
