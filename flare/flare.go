package flare

import 	"github.com/Treeptik/docker-viz/dockertype"

type(
	Flare interface {
		DendrogamAndBubbleImages() string;
		BubbleContainers() string;
	}
)

const(
	BeginJson = "{\"name\": \"Docker\", \"children\": ["
	EndJson = "]}"
)

// Create the json architecture with father/son table
func MakeJsonFather(name string, dockerFils map[string][]string, dockerList map[string]dockertype.DockerType) string {
	var flare string
	nbFils := len(dockerFils[name])
	var i int = 0
	for _, fils := range dockerFils[name] {
		i++
		if _, ok := dockerFils[fils]; ok {
			flare += "{\"name\": \"" + dockerList[fils].GetName() + "\", \"children\": ["
			flare += MakeJsonFather(fils, dockerFils, dockerList) + "]}"
		} else {
			flare += "{\"name\": \"" + dockerList[fils].GetName() + "\", \"size\": " + dockerList[fils].GetSize() + "}"
		}
		if i < nbFils {
			flare += ", "
		}
	}

	return flare
}