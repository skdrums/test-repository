package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// read lines
	reader := bufio.NewReader(file)
	sl := make([]string, 0)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		sl = append(sl, string(line))
	}

	al := Format(sl)
	for _, a := range al {
		fmt.Printf("%#v\n\n", a)
	}
	content0 := al[0]["contents"].(map[string]interface{})
	content2 := al[2]["contents"].(map[string]interface{})
	fmt.Println(len(content0["accounthold_type"].([]interface{})))
	fmt.Println(reflect.TypeOf(content0["accounthold_type"]).Kind() == reflect.Slice)
	fmt.Println(reflect.TypeOf(content2["accounthold_type"]).Kind() == reflect.Slice)

	var str string
	var str2 []string
	for _, con := range content0["accounthold_type"].([]interface{}) {
		str = str + con.(string)
	}

	for _, con := range content2["accounthold_type"].([]interface{}) {
		str2 = append(str2, con.(string))
	}
	fmt.Println(str)
	fmt.Println(strings.Join(str2, ","))

	fmt.Println(reflect.TypeOf(al[0]["client"]))
	client := al[0]["client"].(map[string]interface{})
	fmt.Println(client["ip"])

	var slice []string
	fmt.Println(slice)
}

type ActionLog struct {
	Schema   string    `json:"schema"`
	UUID     string    `json:"uuid"`
	Time     time.Time `json:"time"`
	MineID   string    `json:"mine_id"`
	Platform string    `json:"platform"`
	Client   struct {
		DeviceModel string `json:"device_model"`
		SessionID   string `json:"session_id"`
		Os          string `json:"os"`
		OsVersion   string `json:"os_version"`
		IP          string `json:"ip"`
		Useragent   string `json:"useragent"`
	} `json:"client"`
	ActionType string `json:"action_type"`
	User       struct {
		AsID          string `json:"as_id"`
		AmebaID       string `json:"ameba_id"`
		ServiceUserID string `json:"service_user_id"`
		ActiveUser    bool   `json:"active_user"`
	} `json:"user"`
	Contents map[string]interface{} `json:"contents"`
	Page     struct {
		PageID   string `json:"page_id"`
		Referrer string `json:"referrer"`
	} `json:"page"`
	LogType  string `json:"log_type"`
	LogLevel string `json:"log_level"`
	Env      string `json:"env"`
	NodeURL  string `json:"node_url"`
	LogID    string `json:"log_id"`
	Hostname string `json:"hostname"`
}

func Format(lines []string) []map[string]interface{} {
	al := make([]map[string]interface{}, len(lines))
	for i, line := range lines {
		num := strings.Index(line, "{")
		js := line[num:]
		if err := json.Unmarshal([]byte(js), &al[i]); err != nil {
			fmt.Println(err)
		}
	}
	return al
}

func getStringDate(t time.Time) string {
	year, month, day := t.Date()
	date := strconv.Itoa(year) + strconv.Itoa(int(month)) + strconv.Itoa(day)
	return date
}
