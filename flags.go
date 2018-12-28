package main

import (
	"flag"
	"fmt"
	"strings"
)

// TagFlags is the custom interface to get tags from command line
type TagFlags []string

// String is part of the Value interface of flags.
// used to format the flag's value
func (tagFlags *TagFlags) String() string {
	return fmt.Sprint(*tagFlags)
}

// Set is part of the value interface of flags
// used to set the value from cmd to the interface in code
func (tagFlags *TagFlags) Set(value string) error {
	for _, val := range strings.Split(value, ";") {
		*tagFlags = append(*tagFlags, val)
	}

	return nil
}

// UserData is the model we use to represent the user commnad line input, instead of using separate variables
type UserData struct {
	titleFlag *string
	descFlag  *string
	tagFlags  TagFlags
	allArgs   []string
}

func (userData UserData) String() string {
	return fmt.Sprintf("%+v\n%+v\n%+v\n%+v\n%+v", *(userData.titleFlag), *(userData.descFlag), userData.tagFlags, len(userData.allArgs), userData.allArgs)
}

var userData UserData

// InitFlags inits the flag with required conf
func InitFlags() {
	userData.titleFlag = flag.String("title", "", "The title of the event")
	userData.descFlag = flag.String("desc", "", "The actual event description")
	flag.Var(&userData.tagFlags, "tags", "The tags for the event")

	flag.Parse()
	userData.allArgs = flag.Args()
}
