// package rule provides tools for writing rules for a Hue bridge
package rule

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

var ErrDecode = errors.New("decode error")

type Status string

const (
	Enabled Status = "enabled"
)

type Rule struct {
	Name           string       `json:"name"`
	LastTriggered  Timestamp    `json:"lasttriggered"`
	CreationTime   Timestamp    `json:"creationtime"`
	TimesTriggered int          `json:"timetriggered"`
	Owner          string       `json:"owner"`
	Status         Status       `json:"status"`
	Conditions     []*Condition `json:"conditions"`
	Actions        []*Action    `json:"actions"`
}

type Rules map[string]*Rule
