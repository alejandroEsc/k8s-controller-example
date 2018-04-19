package cloudProviders

func getProvider(name string) CloudProvider {
	switch name {
	case "aws":
		return nil
	}

	return nil
}

// CloudProvider is and interface to be used by structs that are providers
type CloudProvider interface {
	name() string
}
