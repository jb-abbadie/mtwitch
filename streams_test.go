package main

import (
	"errors"
	"testing"
	"time"

	twitch "github.com/grsakea/go-twitch"
)

type fakeTwitchGetter struct {
}

func (s fakeTwitchGetter) GetStream(input twitch.GetStreamInput) (twitch.StreamList, error) {
	if input.UserLogin == "twitch" {
		sl := twitch.StreamList{Data: []twitch.Stream{{}}}
		return sl, nil
	} else {
		return twitch.StreamList{}, errors.New("test")
	}
}

func TestIsOnline(t *testing.T) {
	s := fakeTwitchGetter{}
	out, err := isOnline("twitch", s)
	if err != nil {
		t.Fail()
	}
	if !out {
		t.Fail()
	}
}

func TestIsOnlineFail(t *testing.T) {
	s := fakeTwitchGetter{}
	_, err := isOnline("not_twitch", s)
	if err == nil {
		t.Fail()
	}
}

func TestStreamFilename(t *testing.T) {
	s := twitch.Stream{Title: "@fake_stream!|"}
	tim, _ := time.Parse(time.RFC3339, "2017-01-01T15:04:05Z")
	out := streamFilename(s, tim)
	if out != "17-01-01_15:04-_fake_stream_.mp4" {
		t.Fail()
	}
}

func TestRecordStream(t *testing.T) {

}
