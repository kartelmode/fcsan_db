package config

import (
	"encoding/json"
	"fmt"
	"main/apperrors"
	"os"
)

type Database struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

func CheckDatabaseFields(database *Database) error {
	if database.User == "" {
		return &apperrors.ErrInternal{
			Message: "invalid database user",
		}
	}
	if database.DBName == "" {
		return &apperrors.ErrInternal{
			Message: "invalid database name",
		}
	}
	return nil
}

func ReadConfig(path string, cfg *Database) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return &apperrors.ErrInternal{
			Message: fmt.Sprintf("cannot read config file: %v", err.Error()),
		}
	}

	err = json.Unmarshal(data, cfg)
	if err != nil {
		return &apperrors.ErrInternal{
			Message: fmt.Sprintf("cannot parse config file: %v", err.Error())}
	}

	err = CheckDatabaseFields(cfg)
	if err != nil {
		return &apperrors.ErrInternal{
			Message: fmt.Sprintf("Invalid config field: %v", err.Error())}
	}

	return nil
}
