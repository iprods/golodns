package resolver

type Resolve struct {
	Path string
}

type Domain struct {
	Name      string
	IpAddress string
	Port      string
	Managed   bool
}
