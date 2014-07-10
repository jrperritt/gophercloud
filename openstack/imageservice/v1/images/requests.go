package images

import (
	"fmt"
	"github.com/racker/perigee"
	imageservice "github.com/rackspace/gophercloud/openstack/imageservice/v1"
	"github.com/rackspace/gophercloud/openstack/utils"
	"net/http"
)

type GetResult *http.Response
type UpdateResult *http.Response

func List(c *imageservice.Client, opts ListOpts) ([]Image, error) {
	var i []Image
	url := ""
	h, err := c.GetHeaders()
	if err != nil {
		return nil, err
	}

	query := utils.BuildQuery(opts.Params)

	if !opts.Full {
		url = c.GetListURL() + query
	} else {
		url = c.GetListDetailURL() + query
	}

	_, err = perigee.Request("GET", url, perigee.Options{
		Results: &struct {
			Image *[]Image `json:"images"`
		}{&i},
		MoreHeaders: h,
	})
	return i, err
}

func Get(c *imageservice.Client, opts GetOpts) (GetResult, error) {
	h, err := c.GetHeaders()
	if err != nil {
		return nil, err
	}

	url := c.GetDetailURL(opts.Id)

	resp, err := perigee.Request("GET", url, perigee.Options{
		MoreHeaders: h,
	})
	return &resp.HttpResponse, err
}

func Update(c *imageservice.Client, opts UpdateOpts) (UpdateResult, error) {
	var resp *perigee.Response
	h, err := c.GetHeaders()
	if err != nil {
		return nil, err
	}

	for k, v := range opts.Metadata {
		h["x-image-meta-property-"+k] = v
	}

	url := c.GetUpdateURL(opts.Id)
	fmt.Println(url)
	fmt.Println(opts.Body)
	if opts.Body != nil {
		resp, err = perigee.Request("POST", url, perigee.Options{
			MoreHeaders: h,
			ReqBody:     opts.Body,
		})
	} else {
		resp, err = perigee.Request("POST", url, perigee.Options{
			MoreHeaders: h,
		})
	}
	return &resp.HttpResponse, err
}
