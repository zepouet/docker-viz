package main

import (
	"os"
	"github.com/stretchr/testify/assert"
	"testing"
	"strconv"
)

func init() {

}

func TestLoadConfig(t *testing.T) {
	vizPort, dockerClient := LoadConfig()

	os_vizPort := os.Getenv("VIZ_PORT")
	os_dockerClient := os.Getenv("DOCKER_HOST")

	if os_vizPort == "" {
		os_vizPort = "8080"
	}

	if os_dockerClient == "" {
		os_dockerClient = "unix:///var/run/docker.sock"
	}

	assert.Equal(t, os_dockerClient, dockerClient)
	assert.Equal(t, os_vizPort, strconv.Itoa(vizPort))
}
