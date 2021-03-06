package deploy

import (
	"fmt"

	"github.com/willeponken/ahfdeploy/deploy/aws"
	"github.com/willeponken/ahfdeploy/provider"
)

type Client interface {
	Region() (region string)
	Provider() (providerID int)
	Upload(imageURI string) (result string, err error)
	Create(serviceName, clusterName, containerName string) (result string, err error)
}

func NewClient(providerID int, region string) (client Client, err error) {
	switch providerID {
	case provider.AWS:
		client, err = aws.NewClient(region)
	default:
		err = fmt.Errorf("unknown provider ID: %d", providerID)
	}

	return
}

func NewClientWithCredentials(providerID int, region, id, secret, token string) (client Client, err error) {
	switch providerID {
	case provider.AWS:
		client, err = aws.NewClientWithCredentials(region, id, secret, token)
	default:
		err = fmt.Errorf("unknown provider ID: %d", providerID)
	}

	return
}
