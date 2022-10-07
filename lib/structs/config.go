package structs

type Config struct {
	Port        string `json:"port"`
	AppEnv      string `json:"app_env"`
	MongoURI    string `json:"mongo_uri"`
	DBName      string `json:"db_name"`
	TokenSecret string `json:"token_secret"`
}
