package database

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/denisenkom/go-mssqldb" // SQL driver
// )

// type DBConfig struct {
// 	Host     string
// 	Port     int
// 	User     string
// 	Password string
// 	DBName   string
// 	SSLMode  string
// }

// func NewSQLConnection(cfg DBConfig) (*sql.DB, error) {
// 	connStr := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

// 	// connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
// 	// 	cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

// 	db, err := sql.Open("sqlserver", connStr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = db.Ping(); err != nil {
// 		return nil, err
// 	}

// 	log.Println("Successfully connected to database")
// 	return db, nil
// }
