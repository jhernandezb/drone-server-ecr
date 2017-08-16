package ecr

import (
	"fmt"

	"github.com/awslabs/amazon-ecr-credential-helper/ecr-login/api"
	"github.com/drone/drone/model"
)

type ecrRegistryService struct {
	builtin       model.RegistryService
	ClientFactory api.ClientFactory
}

// Ensure ecrRegistryService implements the interface.
var _ model.RegistryService = &ecrRegistryService{}

// New returns a RegistryService implementation
func New(builtin model.RegistryService) model.RegistryService {
	return &ecrRegistryService{builtin, nil}
}

func (e *ecrRegistryService) RegistryFind(repo *model.Repo, name string) (*model.Registry, error) {
	registry, err := api.ExtractRegistry(name)
	if err != nil {
		// If it is not a valid aws registry url, fallback to the builtin.
		return e.builtin.RegistryFind(repo, name)
	}
	client := e.ClientFactory.NewClientFromRegion(registry.Region)
	auth, err := client.GetCredentials(name)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving creds :%v ", err)
	}
	return &model.Registry{Address: name, Username: auth.Username, Password: auth.Password}, nil
}

func (e *ecrRegistryService) RegistryList(repo *model.Repo) ([]*model.Registry, error) {
	return e.RegistryList(repo)
}

func (e *ecrRegistryService) RegistryCreate(repo *model.Repo, registry *model.Registry) error {
	return e.RegistryCreate(repo, registry)
}

func (e *ecrRegistryService) RegistryUpdate(repo *model.Repo, registry *model.Registry) error {
	return e.RegistryUpdate(repo, registry)
}

func (e *ecrRegistryService) RegistryDelete(repo *model.Repo, name string) error {
	return e.RegistryDelete(repo, name)
}
