package dockertype

type DockerType interface {
	GetId() string
	GetFatherId() string
	GetName() string
	GetSize() string
}