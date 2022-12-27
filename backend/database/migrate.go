package database

import (
	m "case-project/models"
)

// ANCHOR - Migrate database
var (
	migrateModelList = []interface{}{
		&m.User{},
		&m.File{},
	}
)
