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

// IgnitionEngineProperties are the properties that exist in an IgnitionEngine Application
type IgnitionEngineProperties struct {
	Environment       string             `json:"environment"`
	DB                DatabaseProperties `json:"db"`
	ClientInstanceKey string             `json:"client_key"`
	TargetChannel     string             `json:"target_channel"`
	ApplicationName   string             `json:"app_name"`
}
