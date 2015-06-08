package api

// ContainersService provides the containers management aspects of a Docker
// server.
type ContainersService interface {
	List(*ListContainersParams) ([]*Container, error)
}

// ExtendedContainerService provides the containers management aspects of a
// Docker server, including those who cannot be implemented over a strictly
// HTTP/1.x compliant connection.
type ExtendedContainerService interface {
	ContainersService
}

type ListContainersParams struct {
	All     bool
	Limit   int
	Since   int
	Before  int
	Size    bool
	Filters map[string][]string
}

// Port associates a container port to a host port bound to a particular IP
// address.
type Port struct {
	IP          string
	PrivatePort int
	PublicPort  int
	Type        string
}

// Container holds data for an existing container.
type Container struct {
	ID         string `json:"Id"`
	Names      []string
	Image      string
	Command    string
	Created    int
	Ports      []Port
	SizeRw     int `json:",omitempty"`
	SizeRootFs int `json:",omitempty"`
	Labels     map[string]string
	Status     string
}
