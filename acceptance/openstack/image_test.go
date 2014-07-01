// +build acceptance

package openstack

import (
	"fmt"
	"github.com/rackspace/gophercloud/openstack/identity"
	imageservice "github.com/rackspace/gophercloud/openstack/imageservice/v1"
	"github.com/rackspace/gophercloud/openstack/imageservice/v1/images"
	"github.com/rackspace/gophercloud/openstack/utils"
	"os"
	"testing"
)

var serviceType = "image"

func getClient() (*imageservice.Client, error) {
	ao, err := utils.AuthOptions()
	if err != nil {
		return nil, err
	}

	r, err := identity.Authenticate(ao)
	if err != nil {
		return nil, err
	}

	sc, err := identity.GetServiceCatalog(r)
	if err != nil {
		return nil, err
	}

	ces, err := sc.CatalogEntries()
	if err != nil {
		return nil, err
	}

	var eps []identity.Endpoint
	for _, ce := range ces {
		if ce.Type == serviceType {
			eps = ce.Endpoints
		}
	}

	region := os.Getenv("OS_REGION_NAME")
	rep := ""
	for _, ep := range eps {
		if ep.Region == region {
			rep = ep.PublicURL
		}
	}

	client := imageservice.NewClient(rep, r, ao)
	fmt.Printf("%s\n", rep)
	return client, nil

}

func TestImages(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	li, err := images.List(client, images.ListOpts{
		Params: map[string]string{},
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%+v\n", li)
}
