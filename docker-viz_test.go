package main

import (
	"os"
	"github.com/stretchr/testify/assert"
	"testing"
	"strconv"
	"os/exec"
	"log"
)

func init() {
	_, err := exec.Command("./dockerInit.sh").Output()
	if err != nil {
		log.Fatalf("Init: %s", err)
	}

}

func TestLoadConfig(t *testing.T) {
	vizPort := LoadConfig()

	os_vizPort := os.Getenv("VIZ_PORT")

	if os_vizPort == "" {
		os_vizPort = "8080"
	}

	assert.Equal(t, os_vizPort, strconv.Itoa(vizPort))
}
