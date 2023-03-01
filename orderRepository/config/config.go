package config

type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetConfig() *StorageConfig {
	const (
		defHost     = "localhost"
		deftPort    = "5432"
		defDatabase = "postgres"
		defUsername = "postgres"
		defPassword = "342"
	)

	storageConfig := &StorageConfig{
		Host:     defHost,
		Port:     deftPort,
		Database: defDatabase,
		Username: defUsername,
		Password: defPassword,
	}
	return storageConfig
}
