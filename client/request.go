package client

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/isther/translation-tool/config"
)

var (
	uri = "https://openapi.youdao.com/api"
)

type Request struct {
	Q        string `json:"q"`
	From     string `json:"from"`
	To       string `json:"to"`
	AppKey   string `json:"appKey"`
	Salt     string `json:"salt"`
	Sign     string `json:"sign"`
	SignType string `json:"signType"`
	CurTime  int64  `json:"curtime"`
}

func NewRequest(text string) *Request {
	c := &Request{
		Q:        text,
		From:     config.Instance.From,
		To:       config.Instance.To,
		AppKey:   config.Instance.AppKey,
		SignType: config.Instance.SignType,
		CurTime:  time.Now().Unix(),
	}
	c.getSign()
	return c
}

func (c *Request) getSign() {
	input := ""
	if len(c.Q) > 20 {
		input = fmt.Sprintf("%s%d%s", c.Q[:10], len(c.Q), c.Q[len(c.Q)-10:])
	} else {
		input = c.Q
	}
	salt, _ := uuid.NewRandom()
	c.Salt = salt.String()

	data := fmt.Sprintf("%s%s%s%d%s",
		config.Instance.AppKey,
		input,
		c.Salt,
		c.CurTime,
		config.Instance.SecKey,
	)

	c.Sign = EncodeSHA256(data)
}

func EncodeSHA256(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

func (c *Request) Post() (Response, error) {
	res, err := http.PostForm(uri, url.Values{
		"q":        {c.Q},
		"from":     {c.From},
		"to":       {c.To},
		"salt":     {c.Salt},
		"appKey":   {c.AppKey},
		"sign":     {c.Sign},
		"signType": {c.SignType},
		"curtime":  {fmt.Sprintf("%d", c.CurTime)},
	})
	if err != nil {
		fmt.Println(err.Error())
		return Response{}, nil
	}
	defer res.Body.Close()

	var r Response
	body, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal([]byte(body), &r)
	return r, nil
}
