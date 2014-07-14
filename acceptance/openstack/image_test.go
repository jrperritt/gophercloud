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
	//fmt.Printf("%s\n", rep)
	return client, nil

}

func TestImages(t *testing.T) {
	metadata := map[string]string{
		"gopher": "cloud",
	}

	client, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	li, err := images.List(client, images.ListOpts{
		Params: map[string]string{},
		Full:   true,
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%+v\n", li)

	ai, err := images.Add(client, images.AddOpts{})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%+v\n", ai)

	li, err = images.List(client, images.ListOpts{
		Params: map[string]string{},
		Full:   false,
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%+v\n", li)

	ui, err := images.Update(client, images.UpdateOpts{
		Id:       li[0].Id,
		Metadata: metadata,
	})
	fmt.Printf("\n%+v\n", ui)

	gr, err := images.Get(client, images.GetOpts{
		Id: li[0].Id,
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%+v\n", gr)
	/*
		em := images.ExtractMetadata(gr)
		fmt.Printf("\n%+v\n", em)

		ec, err := images.ExtractContent(gr)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Printf("\n%+v\n", ec[0:300])
	*/
}
