package client

import (
	"log"
	"testing"
)

func TestDaily(t *testing.T) {
	daily := newdaily()
	log.Println(daily)
}
