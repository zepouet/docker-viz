# docker-viz [![GoDoc](https://godoc.org/github.com/Treeptik/docker-viz?status.svg)](https://godoc.org/github.com/Treeptik/docker-viz)
docker-viz is a web server who retun a data visualization for different information on [Docker](http://www.dockers.com) containers and images.

## Install and Lauch
### Download & Compile
```
go get github.com/Treeptik/docker-viz
go build $GOPATH/github.com/Treeptik/docker-viz/docker-viz.go
```

### Configure & Launch
docker-viz using two system variables for its configuration.
```
DOCKER_HOST #default value : unix:///var/run/docker.sock
VIZ_PORT #default value : 8080
```


## Visualization type implemented
### Docker Images
- [x] Tree visualization sort by father/son
- [ ] Buble visualization sort by Parent & virtual size

### Docker Containers
- [ ] Buble visualization sort by images & size,ram,uptime
