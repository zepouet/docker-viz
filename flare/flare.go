package flare

import (
	"github.com/Treeptik/docker-viz/dockertype"
	"fmt"
)

const(
	BeginJson = "{\"name\": \"Docker\", \"children\": ["
	EndJson = "]}"
)

// Create the json architecture with father/son table
func MakeJsonFather(name string, dockerFils map[string][]string, dockerList map[string]dockertype.DockerType) string {
	var flare string
	var i int = 0
	for _, fils := range dockerFils[name] {
		if _, ok := dockerFils[fils]; ok {
			if i != 0 {
				flare += ", "
			}
			flare += "{\"name\": \"" + dockerList[fils].GetName() + "\", \"children\": ["
			flare += MakeJsonFather(fils, dockerFils, dockerList) + "]}"
		} else {
			if dockerList[fils].GetSize() == "0" {
				continue
			}
			if i != 0 {
				flare += ", "
			}
			flare += "{\"name\": \"" + dockerList[fils].GetName() + "\", \"size\": " + dockerList[fils].GetSize() + "}"
		}
		i++
	}

	return flare
}