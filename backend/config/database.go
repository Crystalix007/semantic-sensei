package config

import "fmt"

// Database represents the configuration for the database connection.
type Database struct {
	DB       string `yaml:"db"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	SSL      string `yaml:"ssl"`
	User     string `yaml:"user"`
}

// ConnectionString returns the connection string for the database.
func (d Database) ConnectionString() string {
	return fmt.Sprintf(
		"dbname=%s host=%s password=%s sslmode=%s user=%s",
		d.DB,
		d.Host,
		d.Password,
		d.SSL,
		d.User,
	)
}
