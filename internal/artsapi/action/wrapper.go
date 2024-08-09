package action

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi"
)

type ClientWrapper struct {
	Client *artsapi.Client
}

func NewClientWrapper(client *artsapi.Client) (*ClientWrapper, error) {
	if client.Character == "" {
		return nil, fmt.Errorf("character select is required")
	}

	return &ClientWrapper{Client: client}, nil
}
