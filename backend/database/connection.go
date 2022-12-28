package database

import (
	"fmt"
	"go-typescript/secret"

	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

// Anchors - Credentials for database connection and schemas
type Credentials struct {
	host    string
	port    string
	user    string
	pass    string
	sslMode string
	dBName  string
	schemas []string
}

// Anchors - Database struct for database connection
type Database struct {
	myCredentials Credentials
	DB            *gorm.DB
}

var (
	Conn *Database
)

// Anchors - Connect to database with credentials from .env file
func (db *Database) Connect(dbname ...string) {
	//ANCHOR - Close connection if already connected
	if Conn != nil {
		dbCon, _ := db.DB.DB()
		dbCon.Close()
	}
	//ANCHOR - Get credentials from .env file
	envPostEnv := secret.Env["myCredentials"].(map[string]interface{})
	db = &Database{}
	db.myCredentials = Credentials{
		host:    envPostEnv["host"].(string),
		port:    envPostEnv["port"].(string),
		user:    envPostEnv["user"].(string),
		pass:    envPostEnv["pass"].(string),
		sslMode: envPostEnv["sslMode"].(string),
		dBName:  envPostEnv["dbName"].(string),
	}
	//ANCHOR - Create schemas if not exists in database (for development)
	if len(dbname) != 0 {
		db.myCredentials.dBName = dbname[0]
	}
	for _, v := range envPostEnv["schemas"].([]interface{}) {
		db.myCredentials.schemas = append(db.myCredentials.schemas, v.(string))
	}
	dsn := "host=" + db.myCredentials.host + " port=" + db.myCredentials.port + " user=" + db.myCredentials.user + " password=" + db.myCredentials.pass + " dbname=" + db.myCredentials.dBName + " sslmode=" + db.myCredentials.sslMode

	dba, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
	db.DB = dba
	Conn = db
}

func (db *Database) Seed() error {
	var err error
	// dbs := db.DB
	//DESC - All Seeds are created down in order
	// models.User{}.Seed(dbs)
	//DESC - All Seeds are called down in order

	return err
}

// !ANCHOR - MIGRATE DATABASE SCHEMA AND MODELS TO DATABASE (MIGRATE)
func (db *Database) Migrate() {
	var err error
	//!ANCHOR - Migrate Models to Database (Migrate)
	for {
		for _, schema := range db.myCredentials.schemas {

			//!ANCHOR - Create Schema if not exists (Migrate)
			Conn.DB.Exec("CREATE SCHEMA IF NOT EXISTS " + schema + " AUTHORIZATION " + Conn.myCredentials.user + ";")
			fmt.Println("Migrating schema: " + schema)

			//ANCHOR - Any model that has a field with type *gorm.Model will be migrated
			err = Conn.DB.AutoMigrate(migrateModelList...)
			fmt.Println(err)

		}
		if err != nil {
			fmt.Println(err)
		}
		if err == nil {
			break
		}
	}
}

// !SECTION - ENVIRONMENT VARIABLES
func (a Database) String() string {
	return fmt.Sprintf("%s:%s@%s:%s/%s", a.myCredentials.user, a.myCredentials.pass, a.myCredentials.host, a.myCredentials.port, a.myCredentials.dBName)
}

// ANCHOR - Close database connection
func (db *Database) Close() {
	sqlDB, _ := db.DB.DB()
	sqlDB.Close()
}

// ANCHOR - Drop Schema if exists (Migrate)
func (db *Database) DropSchema(schema string) {
	if err := db.DB.Exec("DROP SCHEMA IF EXISTS " + schema + " CASCADE;").Error; err != nil {
		panic(err)
	}
}
