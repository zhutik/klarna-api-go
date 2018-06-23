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
