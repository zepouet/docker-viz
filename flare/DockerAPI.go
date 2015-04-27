package flare

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
		// if docker connection fail, test if user use boot2docker
		_, boot2dockerErr := os.Open(os.Getenv("HOME") + "/.boot2docker")
		if boot2dockerErr != nil {
			// if don't use boot2docker, display the first connection error
			log.Fatal(err)
		} else {
			// else, user use boot2docker. Generate a certificate for boot2docker connection
			caFile := os.Getenv("HOME") + "/.boot2docker/certs/boot2docker-vm/ca.pem"
			certFile := os.Getenv("HOME") + "/.boot2docker/certs/boot2docker-vm/cert.pem"
			keyFile := os.Getenv("HOME") + "/.boot2docker/certs/boot2docker-vm/key.pem"

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