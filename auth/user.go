package authorization

type Login struct {
	GUID          string `json:"username"`
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}
