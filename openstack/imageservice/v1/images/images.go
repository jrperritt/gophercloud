package images

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
}

type ListOpts struct {
	Params map[string]string
}
