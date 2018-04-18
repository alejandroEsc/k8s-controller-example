package cloudProviders



func getProvider(name string) CloudProvider {
	switch (name) {
		case "aws":
			return nil
	}

	return nil
}


type CloudProvider interface {
	name() string
}