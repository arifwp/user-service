package config

var Config AppConfig

type AppConfig struct {
	Port                   int      `json:"port"`
	AppName                string   `json:"appName"`
	AppEnv                 string   `json:"appEnv"`
	SignatureKey           string   `json:"signatureKey"`
	Database               Database `json:"database"`
	RateLimiterMaxRequests int      `json:"rateLimiterMaxRequests"`
	RateLimiterTimeSeconds int      `json:"rateLimiterTimeSeconds"`
	JwtSecretKey           string   `json:"jwtSecretKey"`
	JwtExpirationTime      int      `json:"jwtExpirationTime"`
}

type Database struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxOpenConnection     int    `json:"maxOpenConnection"`
	MaxLifeTimeConnection int    `json:"maxLifeTimeConnection"`
	MaxIdleConnection     int    `json:"maxIdleConnection"`
	MaxIdleTime           int    `json:"maxIdleTime"`
}
