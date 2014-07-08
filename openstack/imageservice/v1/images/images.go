package images

import (
	"io/ioutil"
	"strings"
)

type Image struct {
	Status           string
	Container_format string
	Created_at       string
	Deleted_at       string
	Properties       map[string]interface{}
	Is_public        bool
	Size             int
	Name             string
	Deleted          bool
	Disk_format      string
	Updated_at       string
	Id               string
	Checksum         string
	Owner            string
	Min_disk         int
	Protected        bool
	Min_ram          int
	Location         string
}

type ListOpts struct {
	Full   bool
	Params map[string]string
}

type GetOpts struct {
	Id string
}

func ExtractContent(gr GetResult) ([]byte, error) {
	var body []byte
	defer gr.Body.Close()
	body, err := ioutil.ReadAll(gr.Body)
	return body, err
}

func ExtractMetadata(gr GetResult) map[string]string {
	metadata := make(map[string]string)
	for k, v := range gr.Header {
		if strings.HasPrefix(k, "X-Image-Meta-") {
			key := strings.TrimPrefix(k, "X-Image-Meta-")
			metadata[key] = v[0]
		}
	}
	return metadata
}
