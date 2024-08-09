package database

import (
	"errors"
	"log"
)

// Registering model to database as a migration candidate
func (cp *ConnectionPool) Register(model interface{}) error {
	isRegistered := false
	for _, registered := range cp.models {
		if registered == model {
			isRegistered = true
		}
	}

	if isRegistered {
		return errors.New("Model is already registered")
	}

	cp.models = append(cp.models, model)

	return nil
}

// Migrate and sync all registered model to database
func (cp *ConnectionPool) Migrate() error {
	if len(cp.models) < 1 {
		log.Println("No model registered")
		return nil
	}

	err := cp.DB.AutoMigrate(cp.models...)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
