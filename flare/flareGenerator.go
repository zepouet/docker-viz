package flare

import (
	"github.com/samalba/dockerclient"
	"log"
)

type DockerImage struct {
	Id string
	ParentId string
	Name string
	Create int64
	Size int64
	VirtualSize int64
}

type Flare interface {
	Dendrogam() string;
}

func GenerateDockerImageList(dockerClient string) map[string]DockerImage {
	docker, _ := dockerclient.NewDockerClient(dockerClient, nil)

	containers, err := docker.ListImages()
	if err != nil {
		log.Fatal(err)
	}

	images := make(map[string]DockerImage)
	for _, c := range containers {
		images[c.Id] = DockerImage{c.Id, c.ParentId, c.RepoTags[0], c.Created, c.Size, c.VirtualSize}
	}

	return images
}

func MakeJson(imagesFilsList []string, dockerImagesFils map[string][]string, dockerImagesList map[string]DockerImage) string {
	var flare string
	nbFils := len(imagesFilsList)
	var i int = 0
	for _, image := range imagesFilsList {
		i++
		if _, ok := dockerImagesFils[image]; ok {
			flare += "{\"name\": \"" + dockerImagesList[image].Name + "\", \"children\": ["
			flare += MakeJson(dockerImagesFils[image], dockerImagesFils, dockerImagesList) + "]}"
		} else {
			flare += "{\"name\": \"" + dockerImagesList[image].Name + "\"}"
		}
		if i < nbFils {
			flare += ", "
		}
	}

	return flare
}

func Dendrogam(dockerClient string) string{
	dockerImagesList := GenerateDockerImageList(dockerClient)
	dockerImagesFils := make(map[string][]string)

	for _, image := range dockerImagesList {
		if _, ok := dockerImagesList[image.ParentId]; ok {
			dockerImagesFils[image.ParentId] = append(dockerImagesFils[image.ParentId], image.Id)
		} else {
			dockerImagesFils["Docker"] = append(dockerImagesFils["Docker"], image.Id)
		}
	}
	return  "{\"name\": \"Docker\", \"children\": [" + MakeJson(dockerImagesFils["Docker"], dockerImagesFils, dockerImagesList) + "]}"
}