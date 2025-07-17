package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	TraceOpsDir = ".trace"
	ConfigFile  = "config.json"
)

// Structs
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Repo struct {
	Repository string `json:"repository"`
	URL        string `json:"url"`
}

type Config struct {
	User        User   `json:"user"`
	Repos       []Repo `json:"repos"`
}

// LoadConfig reads the config from .trace/config.json
func LoadConfig() (*Config, error) {
	configPath := filepath.Join(TraceOpsDir, ConfigFile)

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read config.json: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("invalid JSON format in config.json: %w", err)
	}

	return &cfg, nil
}

// SaveConfig writes the updated config back to file
func SaveConfig(cfg *Config) error {
	configPath := filepath.Join(TraceOpsDir, ConfigFile)

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// AddRepo adds a new repo if it doesn't exist
func AddRepo(newRepo Repo) error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	// Check for duplicate
	for _, repo := range cfg.Repos {
		if repo.Repository == newRepo.Repository {
			return fmt.Errorf("repository '%s' already exists", newRepo.Repository)
		}
	}
	for _,repo := range cfg.Repos{
		if repo.URL == newRepo.URL{
			return fmt.Errorf("repository '%s' already exists", newRepo.Repository)
		}
	}

	cfg.Repos = append(cfg.Repos, newRepo)
	return SaveConfig(cfg)
}

// SetUser updates the name and email in the config
func SetUser(name, email string) error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	cfg.User.Name = name
	cfg.User.Email = email

	return SaveConfig(cfg)
}
