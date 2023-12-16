// config.go
package config

import (
	"fmt"
)

type config struct {
	Database database
}

type database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Get() *config {
	cfg := &config{
		Database: database{
			Host:     "localhost", // Ganti dengan nilai yang sesuai
			Port:     "5432",      // Ganti dengan nilai yang sesuai
			User:     "postgres",  // Ganti dengan nilai yang sesuai
			Password: "1234",      // Ganti dengan nilai yang sesuai
			Name:     "fp2",       // Ganti dengan nilai yang sesuai
		},
	}

	// Print or log the values for verification
	fmt.Printf("Database Host: %s\n", cfg.Database.Host)
	fmt.Printf("Database Port: %s\n", cfg.Database.Port)
	fmt.Printf("Database User: %s\n", cfg.Database.User)
	fmt.Printf("Database Password: %s\n", cfg.Database.Password)
	fmt.Printf("Database Name: %s\n", cfg.Database.Name)

	return cfg
}
