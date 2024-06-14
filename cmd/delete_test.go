package main

import (
	"fmt"
	"testing"
)

func TestGetContainerId(t *testing.T) {
	containerId, err := GetContainerId("docker-logstash-1")
	expected := ""

	if containerId != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, containerId)
		fmt.Println(err)
	}
}
