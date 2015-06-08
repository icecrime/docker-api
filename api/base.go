package api

// BaseService provides the basic functionalities of a Docker server.
type BaseService interface {
	Ping() (string, error)
	Version() (*Version, error)
}

// Version is all the version information (for Docker itself, Go, or the OS)
// for a given Docker service.
type Version struct {
	APIVersion    string `json:"ApiVersion"`
	Arch          string
	GitCommit     string
	GoVersion     string
	OS            string `json:"Os"`
	Version       string
	KernelVersion string `json:",omitempty"`
	Experimental  bool   `json:",omitempty"`
}
