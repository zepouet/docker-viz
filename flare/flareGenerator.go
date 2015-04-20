package flare

import (
	"github.com/samalba/dockerclient"
	"log"
	"strconv"
)

type(
	Flare interface {
		Dendrogam() string;
	}
)

// Load all image information clone and commit in Docker
func GenerateDockerImageList(dockerClient *string) map[string]dockerclient.Image {
	docker, _ := dockerclient.NewDockerClient(*dockerClient, nil)

	containers, err := docker.ListImages()
	if err != nil {
		log.Fatal(err)
	}

	images := make(map[string]dockerclient.Image)
	for _, c := range containers {
		images[c.Id] = *c
	}

	return images
}

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
func MakeJson(imageName string, dockerImagesFils map[string][]string, dockerImagesList map[string]dockerclient.Image) string {
	var flare string
	nbFils := len(dockerImagesFils[imageName])
	var i int = 0
	for _, image := range dockerImagesFils[imageName] {
		i++
		if _, ok := dockerImagesFils[image]; ok {
			flare += "{\"name\": \"" + dockerImagesList[image].RepoTags[0] + "\", \"children\": ["
			flare += MakeJson(image, dockerImagesFils, dockerImagesList) + "]}"
		} else {
			virtualSize := strconv.Itoa(int(dockerImagesList[image].VirtualSize))
			flare += "{\"name\": \"" + dockerImagesList[image].RepoTags[0] + "\", \"size\": " + virtualSize + "}"
		}
		if i < nbFils {
			flare += ", "
		}
	}

	return flare
}

// Returns the full json for Dendrogam diagram
func Dendrogam(dockerClient *string) string {
	dockerImagesList := GenerateDockerImageList(dockerClient)
	dockerImagesChilds := GenerateDockerImageChild(dockerImagesList)

	return  "{\"name\": \"Docker\", \"children\": [" + MakeJson("Docker", dockerImagesChilds, dockerImagesList) + "]}"
}