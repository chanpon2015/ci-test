package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Noby struct {
	Response *Response
	key      string
	endpoint string
	persona  string
	ending   string
}

type Response struct {
	CommandID    string     `json:"commandid"`
	CommandName  string     `json:"commandName"`
	Text         string     `json:"text"`
	Type         string     `json:"type"`
	Mood         float32    `json:"mood"`
	Negaposi     float32    `json:"negaposi"`
	NegaposiList []Negaposi `json:"negaposiList"`
	Emotion      Emotion    `json:"emotion"`
	EmotionList  []Emotion  `json:"emotionList"`
	WordList     []Word     `json:"wordList"`
	Art          string     `json:"art"`
	Org          string     `json:"org"`
	Psn          string     `json:"psn"`
	Loc          string     `json:"loc"`
	Dat          string     `json:"dat"`
	Tim          string     `json:"tim"`
}

type Negaposi struct {
	Score float32 `json:"score"`
	Word  string  `json:"word"`
}

type Emotion struct {
	AngerFear   float32 `json:"angerFear"`
	JoySad      float32 `json:"joysad"`
	LikeDislike float32 `json:"likeDislike"`
	Word        string  `json:"word"`
}

type Word struct {
	Feature string `json:"feature"`
	Start   int    `json:"start"`
	Surface string `json:"surface"`
}

type NobyOption func(*Noby)

func SetPersona(persona string) NobyOption {
	return func(noby *Noby) {
		noby.persona = persona
	}
}

func SetEnding(ending string) NobyOption {
	return func(noby *Noby) {
		noby.ending = ending
	}
}

func NewNoby(key string, options ...NobyOption) *Noby {
	noby := Noby{
		key:      key,
		endpoint: "https://app.cotogoto.ai/webapi/noby.json",
	}
	for _, o := range options {
		o(&noby)
	}
	return &noby
}

// APICall 
func (n *Noby) APICall(text string) error {
	values := url.Values{}
	values.Add("appkey", n.key)
	values.Add("text", text)
	if n.persona != "" {
		values.Add("persona", n.persona)
	}
	if n.ending != "" {
		values.Add("ending", n.ending)
	}
	resp, err := http.Get(n.endpoint + "?" + values.Encode())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var r Response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return err
	}
	n.Response = &r
	return nil
}

func (n *Noby) Chat(text string) (string, error) {
	if err := n.APICall(text); err != nil {
		return "", err
	}
	return n.Response.Text, nil
}
