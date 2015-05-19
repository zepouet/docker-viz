package flare

import (
	"github.com/Treeptik/docker-viz/dockertype"
)

const(
	BeginJson = "{\"name\": \"Docker\", \"children\": ["
	EndJson = "]}"
)

// Create the json architecture with father/son table
func MakeJsonFather(name string, dockerFils map[string][]string, dockerList map[string]dockertype.DockerType) string {
	var flare string
	var i int = 0
	// for all son of "name" father
	for _, fils := range dockerFils[name] {
		// if not a first
		if i != 0 {
			flare += ", "
		}
		// if have a children
		if _, ok := dockerFils[fils]; ok {
			flare += "{\"name\": \"" + dockerList[fils].GetName() + "\", \"children\": ["
			flare += MakeJsonFather(fils, dockerFils, dockerList) + "]}"
		} else {
			if dockerList[fils].GetSize() == "0" {
				continue
			}
			// define size values
			size := dockerList[fils].GetSize()
			flare += "{\"name\": \"" + dockerList[fils].GetName() + "\", \"size\": " + size + "}"
		}
		i++
	}

	return flare
}