package server

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/icecrime/docker-api/api"
)

// baseServer implements BaseService interface by exposing HTTP routes and
// forwarding requests to an underlying implementation.
type baseServer struct {
	*restful.WebService
	impl api.BaseService
}

func newBaseServer(impl api.BaseService) *baseServer {
	s := &baseServer{
		impl:       impl,
		WebService: &restful.WebService{},
	}
	s.installRoutes(s.WebService)
	return s
}

func (s *baseServer) installRoutes(ws *restful.WebService) {
	// Swagger doesn't pick up documentation for API endpoints at the root (it
	// basically mixes up listings and delcarations, and there doesn't seem to
	// be possible to distinguish from the URL).
	ws.Path("/").
		Doc("Common operations")

	ws.Route(ws.GET("_ping").To(s.Ping).
		Doc("Ping the Docker server").
		Returns(200, "Constant answer", ""))

	ws.Route(ws.GET("version").To(s.Version).
		Doc("Show the Docker server versions information").
		Produces(restful.MIME_JSON).
		Returns(200, "Version information", api.Version{}))
}

func (s *baseServer) Ping(request *restful.Request, response *restful.Response) {
	p, err := s.impl.Ping()
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	response.Write([]byte(p))
}

func (s *baseServer) Version(request *restful.Request, response *restful.Response) {
	v, err := s.impl.Version()
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	response.WriteEntity(v)
}
