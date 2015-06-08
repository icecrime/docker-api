package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/icecrime/api/api"
)

func NewContainersServiceClient(client *http.Client, baseURI string) *containersClient {
	return &containersClient{
		baseURI: baseURI,
		client:  client,
	}
}

// containersClient provides client-side implementation of the BaseService interface.
type containersClient struct {
	baseURI string
	client  *http.Client
}

func (b *containersClient) List(_ *api.ListContainersParams) ([]*api.Container, error) {
	r, err := b.client.Get(fmt.Sprintf("%s/containers/ps", b.baseURI))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var out []*api.Container
	if err := json.NewDecoder(r.Body).Decode(&out); err != nil {
		return nil, err
	}

	return out, nil
}
