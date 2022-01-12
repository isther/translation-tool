package client

import (
	"github.com/fatih/color"
)

type Response struct {
	ErrorCode   string   `json:"errorCode"`
	Query       string   `json:"query"`
	Translation []string `json:"translation"`
	Basic       string   `json:"basic"`
	Web         []web    `json:"web"`
}

type web struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

func (r *Response) PrintQuery() {
	color.Green("English: ")
	color.Blue("%v", r.Query)
}

func (r *Response) PrintTranslate() {
	color.Green("翻译结果: \n")
	for i, v := range r.Translation {
		if i > 10 {
			return
		}
		color.Blue("%d: %v", i+1, v)
	}
}

func (r *Response) PrintWeb() {
	color.Green("网络释义: \n")
	for i, v := range r.Web {
		if i > 10 {
			return
		}
		color.Blue("	%v: %v\n", v.Key, v.Value)
	}
}
