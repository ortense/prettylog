package prettylog

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	cs "github.com/ortense/consolestyle"
)

type input struct {
	Time    string                 `json:"time"`
	Level   string                 `json:"level"`
	Message string                 `json:"msg"`
	Extra   map[string]interface{} `json:"-"`
}

func parse(s string) *input {
	data := input{}

	err := json.Unmarshal([]byte(s), &data)

	if err != nil {
		return nil
	}

	var additionalFields map[string]interface{}

	err = json.Unmarshal([]byte(s), &additionalFields)

	if err == nil {
		delete(additionalFields, "time")
		delete(additionalFields, "level")
		delete(additionalFields, "msg")
		if len(additionalFields) > 0 {
			data.Extra = additionalFields
		}
	}

	return &data
}

const timeInputLayout = "2006-01-02T15:04:05.999999999-07:00"
const timeOutputLayout = "15:04:05.000"

func formatTime(s string) string {
	date, err := time.Parse(timeInputLayout, s)
	time := ""

	if err == nil {
		time = cs.Bold(date.Format(timeOutputLayout)) + " "
	}

	return time
}

func formatLevel(s string) string {
	level := strings.ToUpper(s)

	switch level {
	case "ERROR":
		level = cs.Bold(cs.Red(level + " 💥 "))
	case "WARN":
		level = cs.Bold(cs.Yellow(level + " 🚧 "))
	case "INFO":
		level = cs.Bold(cs.Cyan(level + " 💡 "))
	case "":
		level = cs.Bold("📄 ")
	default:
		level = cs.Blue(cs.Bold(level + " 🔍 "))
	}

	return level
}

var jsonProp = regexp.MustCompile(`"([^"]+)":`)
var strValue = regexp.MustCompile(`: "([^"]+)"`)
var boolValue = regexp.MustCompile(`: (true|false)`)

func formatExtra(data map[string]interface{}) string {
	extraJSON, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return ""
	}

	extra := string(extraJSON)
	extra = jsonProp.ReplaceAllString(extra, cs.Green(`"$1"`)+":")
	extra = strValue.ReplaceAllString(extra, ": "+cs.Magenta(`"$1"`))
	extra = boolValue.ReplaceAllString(extra, ": "+cs.Blue(`$1`))

	return extra
}

func Print(s string) {
	data := parse(s)

	if data == nil {
		fmt.Println(s)
		return
	}

	time := formatTime(data.Time)
	level := formatLevel(data.Level)

	fmt.Print(time, level, data.Message, " ")

	if len(data.Extra) > 0 {
		extra := formatExtra(data.Extra)

		fmt.Println(extra)
	}

	fmt.Println("")
}
