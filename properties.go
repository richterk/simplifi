package simplifiFramework

type DatabaseProperties struct {
	Host     string `json:"host"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Properties for the API
type Properties struct {
	Environment    string             `json:"environment"`
	DB             DatabaseProperties `json:"db"`
	Port           int                `json:"port"`
	JwtSecret      string             `json:"jwtSecret"`
	GithubAPIToken string             `json:"githubAPIToken"`
}
