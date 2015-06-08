package server

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/icecrime/api/api"
)

// containersServer implements ContainersService by exposing HTTP routes and
// forwarding requests to an underlying implementation.
type containersServer struct {
	*restful.WebService
	impl api.ContainersService
}

func newContainersServer(impl api.ContainersService) *containersServer {
	s := &containersServer{
		impl:       impl,
		WebService: &restful.WebService{},
	}
	s.installRoutes(s.WebService)
	return s
}

func (s *containersServer) installRoutes(ws *restful.WebService) {
	ws.Path("/containers").
		Doc("Containers management").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("ps").To(s.List).
		Doc("List containers").
		Param(ws.QueryParameter("all", "Show all containers").DataType("string").DefaultValue("false")).
		Param(ws.QueryParameter("limit", "Maximum returns (0: unlimited)").DataType("int").DefaultValue("0")).
		Param(ws.QueryParameter("since", "Only show containers created after timestamp").DataType("int").DefaultValue("0")).
		Param(ws.QueryParameter("before", "Only show containers created before timestamp").DataType("int").DefaultValue("0")).
		Param(ws.QueryParameter("size", "Return the containers size").DataType("string").DefaultValue("false")).
		Param(ws.QueryParameter("filters", "Filter containers").DataType("map[string][]string")).
		Returns(200, "Container list", []*api.Container{}))
}

func (s *containersServer) List(request *restful.Request, response *restful.Response) {
	params := &api.ListContainersParams{}

	if all, err := booleanValue(request.QueryParameter("all"), false); err == nil {
		params.All = all
	} else {
		response.WriteError(http.StatusBadRequest, err)
		return
	}

	containerList, err := s.impl.List(params)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	response.WriteEntity(containerList)
}
