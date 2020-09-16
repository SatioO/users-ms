package config

// Configurations exported
type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
	Security SecurityConfiguration
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port string
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
}

// SecurityConfiguration exported
type SecurityConfiguration struct {
	AccessSecret  string
	RefreshSecret string
}
