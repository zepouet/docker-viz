package dockertype

import (
	"github.com/fsouza/go-dockerclient"
	"log"
	"os"
	"fmt"
)

// Load docker env variables
func LoadConfig() string {
	dockerClient := os.Getenv("DOCKER_HOST")

	// if var not defined, change for defaults values
	if dockerClient == "" {
		dockerClient = "unix:///var/run/docker.sock"
	}

	return dockerClient
}

// Create a Docker engine connection
func DockerEngineConnection() *docker.Client {
	dockerConnection, err := docker.NewClient(LoadConfig())
	if err != nil {
		log.Print(err)
	}

	_, err = dockerConnection.Version()
	// if docker api can't find the docker version
	if err != nil {
		// if docker connection fail, test if user use boot2docker
		_, boot2dockerErr := os.Open(os.Getenv("HOME") + "/.boot2docker")
		if boot2dockerErr != nil {
			// if don't use boot2docker, display the first connection error
			log.Print("boot2docker not detected\n")
			log.Print(err)
		} else {
			// else, user use boot2docker. Generate a certificate for boot2docker connection
			path := os.Getenv("DOCKER_CERT_PATH")
			ca := fmt.Sprintf("%s/ca.pem", path)
			cert := fmt.Sprintf("%s/cert.pem", path)
			key := fmt.Sprintf("%s/key.pem", path)

			// Create a boot2docker connection
			dockerConnection, err := docker.NewTLSClient(LoadConfig(), cert, key, ca)

			if err != nil {
				// if connection fail, display a boot2docker connection error
				log.Print(err)
			}

			return dockerConnection
		}
	}

	return dockerConnection
}

// Docker Version
func DockerVersion() string {
	dockerConnection := DockerEngineConnection()

	_, err := dockerConnection.Version()
	if err != nil {
		return "Docker Engine not found"
	}

	return "0.0"
}

// Docker Statut
func DockerStatut() bool {
	dockerConnection := DockerEngineConnection()

	_, err := dockerConnection.Version()
	if err != nil {
		return false
	}

	return true
}

// load images list
func LoadDockerImages() []docker.APIImages {
	dockerConnection := DockerEngineConnection()

	images, err := dockerConnection.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		log.Fatal(err)
	}

	return images
}

// load containers list
func LoadDockerContainers() []docker.APIContainers {
	dockerConnection := DockerEngineConnection()

	containers, err := dockerConnection.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		log.Fatal(err)
	}

	return containers
}

// Load all image information clone and commit in Docker
func GenerateDockerImageList() map[string]DockerType {

	images := make(map[string]DockerType)
	for _, c := range LoadDockerImages() {
		i := Image{c}
		d := DockerType(i)
		images[d.GetId()] = d
	}

	return images
}

func countDockerImages() int {
	return len(GenerateDockerImageList())
}

func countDockerContainer() int {
	return len(LoadDockerContainers())
}

// Load all container information in Docker
func GenerateDockerContainerList() map[string]DockerType {

	container := make(map[string]DockerType)
	for _, c := range LoadDockerContainers() {
		i := Container{c}
		d := DockerType(i)
		container[d.GetId()] = d
	}

	return container
}

// return a map who represent the Father/son sequence
func GenerateDockerChild(dockerList map[string]DockerType) map[string][]string {
	dockerImagesChilds := make(map[string][]string)
	for _, docker := range dockerList {
		if _, ok := dockerList[docker.GetFatherId()]; ok {
			dockerImagesChilds[docker.GetFatherId()] = append(dockerImagesChilds[docker.GetFatherId()], docker.GetId())
		} else {
			dockerImagesChilds["Docker"] = append(dockerImagesChilds["Docker"], docker.GetId())
		}
	}

	return dockerImagesChilds
}

func LoadContainerInfos(Id string) (*docker.Container, bool) {
	dockerConnection := DockerEngineConnection()

	containerInfo, err := dockerConnection.InspectContainer(Id)
	if err != nil {
		return nil, true
	}

	return containerInfo, false
}