package ge

import "github.com/Sysleec/Artifacts-client/internal/artsapi"

type ClientWrapper struct {
	Client *artsapi.Client
}

func NewClientWrapper(client *artsapi.Client) (*ClientWrapper, error) {
	return &ClientWrapper{Client: client}, nil
}
