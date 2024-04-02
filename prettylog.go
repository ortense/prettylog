package prettylog

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	cs "github.com/ortense/consolestyle"
)

type input struct {
	Time  interface{}            `json:"time"`
	Level interface{}            `json:"level"`
	Msg   string                 `json:"msg"`
	Extra map[string]interface{} `json:"-"`
}

type output struct {
	Time    *time.Time
	Level   string
	Message string
	Extra   map[string]interface{}
}

func parseJSON(s string) output {
	input := input{}
	output := output{}

	err := json.Unmarshal([]byte(s), &input)

	if err != nil {
		output.Message = s
		return output
	}

	switch level := input.Level.(type) {
	case string:
		output.Level = level
	case float64:
		output.Level = strconv.FormatFloat(level, 'f', -1, 64)
	}

	switch time := input.Time.(type) {
	case string:
		t, _ := parseTimeString(time)
		output.Time = &t
	case float64:
		t := parseTimestamp(time)
		output.Time = &t
	}

	output.Message = input.Msg

	var additionalFields map[string]interface{}

	err = json.Unmarshal([]byte(s), &additionalFields)

	if err == nil {
		delete(additionalFields, "time")
		delete(additionalFields, "level")
		delete(additionalFields, "msg")
		if len(additionalFields) > 0 {
			output.Extra = additionalFields
		}
	}

	return output
}

const timeStringInputLayout = "2006-01-02T15:04:05.999999999-07:00"
const timeOutputLayout = "15:04:05.000"

func parseTimeString(s string) (time.Time, error) {
	return time.Parse(timeStringInputLayout, s)
}

func parseTimestamp(f float64) time.Time {
	return time.Unix(int64(f/1000), 0)
}

func formatTime(t *time.Time) string {
	if t == nil {
		now := time.Now()
		t = &now
	}

	return cs.Yellow(cs.Bold(t.Format(timeOutputLayout))) + " "
}

func formatLevel(s string) string {
	level := strings.ToUpper(s)

	switch level {
	case "TRACE", "10":
		level = cs.Bold(cs.Red("TRACE 🧶 "))
	case "DEBUG", "20":
		level = cs.Bold(cs.Red("DEBUG 🐞 "))
	case "INFO", "30":
		level = cs.Bold(cs.Cyan("INFO 💡 "))
	case "WARN", "40":
		level = cs.Bold(cs.Yellow("WARN 🚧 "))
	case "ERROR", "50":
		level = cs.Bold(cs.Red("ERROR 💥 "))
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
	data := parseJSON(s)
	time := formatTime(data.Time)
	level := formatLevel(data.Level)

	fmt.Print(time, level, data.Message, " ")

	if len(data.Extra) > 0 {
		extra := formatExtra(data.Extra)

		fmt.Println(extra)
	}

	fmt.Println("")
}
