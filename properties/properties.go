package properties

//DatabaseProperties - This is the object representing the database connection
type DatabaseProperties struct {
	Host     string `json:"host"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

//Amqp - This is the object representing an Amqp connection
type Amqp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

// Properties for the API
type Properties struct {
	Environment    string             `json:"environment"`
	DB             DatabaseProperties `json:"db"`
	Port           int                `json:"port"`
	JwtSecret      string             `json:"jwtSecret"`
	GithubAPIToken string             `json:"githubAPIToken"`
	Amqp           Amqp               `json:"amqp"`
}

// IgnitionEngineProperties are the properties that exist in an IgnitionEngine Application
type IgnitionEngineProperties struct {
	Environment       string             `json:"environment"`
	DB                DatabaseProperties `json:"db"`
	ClientInstanceKey string             `json:"client_key"`
	TargetChannel     string             `json:"target_channel"`
	ApplicationName   string             `json:"app_name"`
}
