package klarna

// Environment allows clients to be configured for Testing
// and Production environments.
type Environment struct {
	apiURL string
}

// Production - instance of production environment
var Production = Environment{
	apiURL: "https://api.klarna.com/",
}

// Testing - instance of testing environment
var Testing = Environment{
	apiURL: "https://api.playground.klarna.com/",
}

// TestEnvironment returns test environment configuration.
func TestEnvironment() Environment {
	return Testing
}

// ProductionEnvironment returns production environment configuration.
func ProductionEnvironment() Environment {
	return Production
}
