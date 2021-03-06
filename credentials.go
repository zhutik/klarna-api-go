package klarna

// apiCredentials basic API settings
//
// Description:
//
//     - Env - Environment for next API calls
//     - Username - API username for authentication
//     - Password - API password for authentication
//
// You can create new API credentials there: https://users.playground.eu.portal.klarna.com/users
type apiCredentials struct {
	Env      Environment
	Username string
	Password string
}

// makeCredentials create new APICredentials
func makeCredentials(env Environment, username, password string) apiCredentials {
	return apiCredentials{
		Env:      env,
		Username: username,
		Password: password,
	}
}
