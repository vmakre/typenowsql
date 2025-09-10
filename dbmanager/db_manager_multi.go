package dbmanager

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/denisenkom/go-mssqldb" // SQL driver
	//	_ "github.com/lib/pq" // PostgreSQL driver
	// _ "github.com/go-sql-driver/mysql" // MySQL driver
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type DBType string

const (
	PostgreSQL DBType = "postgres"
	MySQL      DBType = "mysql"
	SQLite     DBType = "sqlite3"
	Sqlserver  DBType = "sqlserver"
)

type DBConfig struct {
	Type     DBType
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string // For PostgreSQL
}

type DatabaseManager struct {
	connections map[string]*sql.DB
	mu          sync.RWMutex
	configs     map[string]DBConfig
}

func NewDatabaseManager() *DatabaseManager {
	return &DatabaseManager{
		connections: make(map[string]*sql.DB),
		configs:     make(map[string]DBConfig),
	}
}

// AddConnection adds a new database connection configuration
func (dm *DatabaseManager) AddConnection(name string, config DBConfig) error {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	if _, exists := dm.configs[name]; exists {
		return fmt.Errorf("connection '%s' already exists", name)
	}

	dm.configs[name] = config
	return nil
}

// Connect establishes a connection to the database
func (dm *DatabaseManager) Connect(name string) error {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	config, exists := dm.configs[name]
	if !exists {
		return fmt.Errorf("connection '%s' not found", name)
	}

	if _, exists := dm.connections[name]; exists {
		return fmt.Errorf("connection '%s' already established", name)
	}

	connString := dm.buildConnectionString(config)

	db, err := sql.Open(string(config.Type), connString)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return fmt.Errorf("failed to ping database: %w", err)
	}

	dm.connections[name] = db
	log.Printf("Successfully connected to database: %s", name)
	return nil
}

// buildConnectionString constructs the appropriate connection string
func (dm *DatabaseManager) buildConnectionString(config DBConfig) string {
	switch config.Type {
	case PostgreSQL:
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	case MySQL:
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			config.User, config.Password, config.Host, config.Port, config.DBName)
	case SQLite:
		return config.DBName // For SQLite, DBName is the file path
	case Sqlserver:
		return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			config.User, config.Password, config.Host, config.Port, config.DBName)
	default:
		return ""
	}
}

// GetConnection returns a database connection by name
func (dm *DatabaseManager) GetConnection(name string) (*sql.DB, error) {
	dm.mu.RLock()
	defer dm.mu.RUnlock()

	db, exists := dm.connections[name]

	if !exists {
		return nil, fmt.Errorf("connection '%s' not found or not connected", name)
	}

	return db, nil
}

// GetConnection returns a database connection by name
func (dm *DatabaseManager) GetConnections() (map[string]*sql.DB, error) {
	dm.mu.RLock()
	defer dm.mu.RUnlock()

	// for i, col := range dm.connections {
	// 	result[col] = values[i]
	// 		}
	// db, exists :=

	return dm.connections, nil
}

// CloseConnection closes a specific database connection
func (dm *DatabaseManager) CloseConnection(name string) error {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	db, exists := dm.connections[name]
	if !exists {
		return fmt.Errorf("connection '%s' not found", name)
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("failed to close connection '%s': %w", name, err)
	}

	delete(dm.connections, name)
	log.Printf("Closed database connection: %s", name)
	return nil
}

// CloseAll closes all database connections
func (dm *DatabaseManager) CloseAll() {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	for name, db := range dm.connections {
		if err := db.Close(); err != nil {
			log.Printf("Error closing connection '%s': %v", name, err)
		}
	}

	dm.connections = make(map[string]*sql.DB)
	log.Println("All database connections closed")
}

// HealthCheck checks the health of all connections
func (dm *DatabaseManager) HealthCheck() map[string]error {
	dm.mu.RLock()
	defer dm.mu.RUnlock()

	results := make(map[string]error)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for name, db := range dm.connections {
		if err := db.PingContext(ctx); err != nil {
			results[name] = err
		} else {
			results[name] = nil
		}
	}

	return results
}
