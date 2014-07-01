package images

import (
	"github.com/racker/perigee"
	imageservice "github.com/rackspace/gophercloud/openstack/imageservice/v1"
	"github.com/rackspace/gophercloud/openstack/utils"
	//"net/http"
)

//type ListResult *http.Response

func List(c *imageservice.Client, opts ListOpts) ([]Image, error) {
	var i []Image
	h, err := c.GetHeaders()
	if err != nil {
		return nil, err
	}

	query := utils.BuildQuery(opts.Params)

	url := c.GetListURL() + query

	_, err = perigee.Request("GET", url, perigee.Options{
		Results: &struct {
			Image *[]Image `json:"images"`
		}{&i},
		MoreHeaders: h,
	})
	return i, err
}
