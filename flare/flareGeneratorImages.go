package flare

import (
	"github.com/samalba/dockerclient"
	"strconv"
)

// Creating Association father/son
func GenerateDockerImageChild(dockerImagesList map[string]dockerclient.Image) map[string][]string {
	dockerImagesChilds := make(map[string][]string)
	for _, image := range dockerImagesList {
		if _, ok := dockerImagesList[image.ParentId]; ok {
			dockerImagesChilds[image.ParentId] = append(dockerImagesChilds[image.ParentId], image.Id)
		} else {
			dockerImagesChilds["Docker"] = append(dockerImagesChilds["Docker"], image.Id)
		}
	}

	return dockerImagesChilds
}

// Create the json architecture with father/son table
func MakeJsonImages(imageName string, dockerImagesFils map[string][]string, dockerImagesList map[string]dockerclient.Image) string {
	var flare string
	nbFils := len(dockerImagesFils[imageName])
	var i int = 0
	for _, image := range dockerImagesFils[imageName] {
		i++
		if _, ok := dockerImagesFils[image]; ok {
			flare += "{\"name\": \"" + dockerImagesList[image].RepoTags[0] + "\", \"children\": ["
			flare += MakeJsonImages(image, dockerImagesFils, dockerImagesList) + "]}"
		} else {
			virtualSize := strconv.FormatInt(dockerImagesList[image].VirtualSize, 10)
			flare += "{\"name\": \"" + dockerImagesList[image].RepoTags[0] + "\", \"size\": " + virtualSize + "}"
		}
		if i < nbFils {
			flare += ", "
		}
	}

	return flare
}

// Returns the full json for docker images Dendrogam & bubble diagram
func DendrogamAndBubbleImages() string {
	dockerImagesList := GenerateDockerImageList()
	dockerImagesChilds := GenerateDockerImageChild(dockerImagesList)

	return  BeginJson + MakeJsonImages("Docker", dockerImagesChilds, dockerImagesList) + EndJson
}