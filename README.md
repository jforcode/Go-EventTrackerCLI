So, this is a command line tool for event tracker, i.e. you can track events from the command line.

# Usage:

EventTracker-cmd [-title="<TITLE>"] [-desc="<DESC>"] [-tags="<TAGS>"] ["<ARGS>"]
- TITLE is the title of event
- DESC is the description of the event
- TAGS is a <;> separated list of tags for the event
- if -title FLAG is not present, the first available ARG is considered as title
- if -desc FLAG is not present, the first available (after title resolution) is considered as desc
- all other ARGs are considered as <;> separated list of tags
- title is the only mandatory param (should be present either in FLAGs or ARGs)

# Command line gotchas:
- anything after the first ARG is considered as an ARG. So, adding a FLAG after ARG will make it be considered as an ARG
- all <"> are stripped from a command.
- <;> is also used as to execute multiple commands at once in terminal.

# Examples

## Example 1
EventTracker-cmd \
-title="some title" \
-desc="Some desc" \
-tags="tag1;tag 2" \
"tag 3" some other tag mixed1

## Corresponding event
{
    "title": "some title",
    "desc": "Some desc",
    "tags": [
        "tag1",
        "tag 2",
        "tag 3",
        "some",
        "other",
        "tag",
        "mixed1"
    ]
}

## Example 2
EventTracker-cmd \
-tags="tag1;tag 2" \
"Some title" \
"Some desc" \
"one more"

## Corresponding event
{
    "title": "Some title",
    "desc": "Some desc",
    "tags": [
        "tag1",
        "tag 2",
        "one more"
    ]
}

## Example 3
./EventTracker-cmd \
mistakenly_entered_value \
-title="Some title" \
-desc="Some desc" \
-tags="tag1;tag2"

## Corresponding event
{
    "title": "mistakenly_entered_value",
    "desc": "-title=Some title",
    "tags": [
        "-desc=Some desc",
        "-tags=tag1",
        "tag2"
    ]
}
