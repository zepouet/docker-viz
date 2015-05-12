package dockertype

import (
	"github.com/samalba/dockerclient"
	"log"
	"os"
	"crypto/tls"
	"io/ioutil"
	"crypto/x509"
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
func DockerEngineConnection() *dockerclient.DockerClient {
	docker, err := dockerclient.NewDockerClient(LoadConfig(), nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = docker.Version()
	// if docker api can't find the docker version
	if err != nil {
		// if docker connection fail, test if user use boot2docker
		_, boot2dockerErr := os.Open(os.Getenv("HOME") + "/.boot2docker")
		if boot2dockerErr != nil {
			// if don't use boot2docker, display the first connection error
			log.Print("boot2docker not detected\n")
			log.Fatal(err)
		} else {
			// else, user use boot2docker. Generate a certificate for boot2docker connection
			caFile := os.Getenv("DOCKER_CERT_PATH") + "/ca.pem"
			certFile := os.Getenv("DOCKER_CERT_PATH") + "/cert.pem"
			keyFile := os.Getenv("DOCKER_CERT_PATH") + "/key.pem"

			cert, _ := tls.LoadX509KeyPair(certFile, keyFile)
			pemCerts, _ := ioutil.ReadFile(caFile)

			tlsConfig := &tls.Config{}
			tlsConfig.RootCAs       = x509.NewCertPool()
			tlsConfig.ClientAuth    = tls.RequireAndVerifyClientCert
			tlsConfig.Certificates  = []tls.Certificate{cert}
			tlsConfig.RootCAs.AppendCertsFromPEM(pemCerts)

			// Create a boot2docker connection
			docker, err := dockerclient.NewDockerClient(LoadConfig(), tlsConfig)

			if err != nil {
				// if connection fail, display a boot2docker connection error
				log.Fatal(err)
			}

			return docker
		}
	}

	return docker
}

// load images list
func LoadDockerImages() []*dockerclient.Image {
	docker := DockerEngineConnection()

	images, err := docker.ListImages()
	if err != nil {
		log.Fatal(err)
	}

	return images
}

// load containers list
func LoadDockerContainers() []dockerclient.Container {
	docker := DockerEngineConnection()

	containers, err := docker.ListContainers(true, true, "")
	if err != nil {
		log.Fatal(err)
	}

	return containers
}

// load all infos on container id
func LoadContainerInfos(id string) *dockerclient.ContainerInfo {
	docker := DockerEngineConnection()

	infos, err := docker.InspectContainer(id)
	if err != nil {
		log.Fatal(err)
	}

	return infos
}

// Load all image information clone and commit in Docker
func GenerateDockerImageList() map[string]DockerType {

	images := make(map[string]DockerType)
	for _, c := range LoadDockerImages() {
		i := Image{*c}
		d := DockerType(i)
		images[d.GetId()] = d
	}

	return images
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