package flare

import (
	"github.com/samalba/dockerclient"
	"strconv"
)

// Create the json architecture
func MakeJsonContainers(dockerContainerList []dockerclient.Container) string {
	var flare string
	var i int = 0
	for _, container := range dockerContainerList {
		virtualSize := strconv.FormatInt(container.SizeRw, 10)
		if virtualSize == "0" {
			continue
		}
		if i != 0 {
			flare += ", "
		}
		flare += "{\"name\": \"" + container.Names[0] + "\", \"size\": " + virtualSize + "}"
		i++
	}

	return flare
}

// Returns the full json for docker containers bubble diagram
func BubbleContainers() string {
	dockerContainerList := LoadDockerContainers()

	return  BeginJson + MakeJsonContainers(dockerContainerList) + EndJson
}