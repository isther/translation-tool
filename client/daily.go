package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	dailyUrl = "http://open.iciba.com/dsapi"
)

type daily struct {
	Content string `json:"content"`
	Note    string `json:"note"`
}

func PrintDaily() {
	daily := *newdaily()
	color.Green("[Daily Sentence]\n")
	color.Blue("%v\n", daily.Content)
	color.Blue("%v\n", daily.Note)
}

func newdaily() *daily {
	resp, err := http.Post(dailyUrl,
		"application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("date=%v-%v-%v", time.Now().Year(), time.Now().Month(), time.Now().Day())))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	var d daily
	json.Unmarshal(body, &d)

	return &d
}
