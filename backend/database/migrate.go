package database

import (
	m "go-typescript/models"
)

// ANCHOR - Migrate database
var (
	migrateModelList = []interface{}{
		&m.User{},
		&m.File{},
	}
)
