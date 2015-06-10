# docker-viz [![GoDoc](https://godoc.org/github.com/Treeptik/docker-viz?status.svg)](https://godoc.org/github.com/Treeptik/docker-viz) [![Build Status](https://travis-ci.org/Treeptik/docker-viz.svg)](https://travis-ci.org/Treeptik/docker-viz)
docker-viz is a web server who return a data visualization for different information on [Docker](http://www.docker.com) containers and images.

## Install and Lauch
### Download & Compile
```
go get github.com/Treeptik/docker-viz
go build $GOPATH/github.com/Treeptik/docker-viz/docker-viz.go
```

### Configure & Launch
Docker-Viz using one system variable for its configuration.
```
DOCKER_HOST #default value : unix:///var/run/docker.sock
```

Docker-Viz start usage
```
Usage: 
  docker-viz [flags]
  docker-viz [command]

Available Commands: 
  version     docker-viz version
  help        Help about any command

Flags:
  -d, --debug=false: Run docker-viz server in "debug" mode
  -h, --help=false: help for docker-viz
  -p, --port=8080: docker-viz server port


Use "docker-viz help [command]" for more information about a command
```


## Visualization type implemented
### Docker Images
- [x] Tree visualization sort by father/son
- [x] Bubble visualization sort by Parent & virtual size

### Docker Containers
- [x] Bubble visualization sort by images & size
- [ ] Chord visualization sort by link, volume
- [x] Matrice visualization by link & volume

### Docker Engine
- [x] General information visualization
