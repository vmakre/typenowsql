package database

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"time"

// 	_ "github.com/denisenkom/go-mssqldb" // SQL driver
// 	//	_ "github.com/lib/pq" // PostgreSQL driver
// 	// _ "github.com/go-sql-driver/mysql" // MySQL driver
// 	// _ "github.com/mattn/go-sqlite3"    // SQLite driver
// )

// type DBManager struct {
// 	dsn          string
// 	db           *sql.DB
// 	mu           sync.RWMutex
// 	initialized  bool
// 	maxOpenConns int
// 	maxIdleConns int
// 	maxLifetime  time.Duration
// }

// func NewDBManager(dsn string, maxOpen, maxIdle int, maxLifetime time.Duration, initialized bool) *DBManager {
// 	return &DBManager{
// 		dsn:          dsn,
// 		maxOpenConns: maxOpen,
// 		maxIdleConns: maxIdle,
// 		maxLifetime:  maxLifetime,
// 		initialized:  initialized,
// 	}
// }

// func (m *DBManager) GetDB() (*sql.DB, error) {
// 	m.mu.RLock()
// 	if m.initialized && m.db != nil {
// 		db := m.db
// 		m.mu.RUnlock()
// 		return db, nil
// 	}
// 	m.mu.RUnlock()

// 	m.mu.Lock()
// 	defer m.mu.Unlock()

// 	// Double-check after acquiring write lock
// 	if m.initialized && m.db != nil {
// 		return m.db, nil
// 	}

// 	// Initialize new connection
// 	db, err := sql.Open("sqlserver", m.dsn)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to open database: %w", err)
// 	}

// 	// Configure connection pool
// 	db.SetMaxOpenConns(m.maxOpenConns)
// 	db.SetMaxIdleConns(m.maxIdleConns)
// 	db.SetConnMaxLifetime(m.maxLifetime)

// 	// Test connection
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	if err := db.PingContext(ctx); err != nil {
// 		db.Close()
// 		return nil, fmt.Errorf("failed to ping database: %w", err)
// 	}

// 	m.db = db
// 	m.initialized = true

// 	return m.db, nil
// }

// func (m *DBManager) Close() error {
// 	m.mu.Lock()
// 	defer m.mu.Unlock()

// 	if m.db != nil {
// 		err := m.db.Close()
// 		m.db = nil
// 		m.initialized = false
// 		return err
// 	}
// 	return nil
// }

// func (m *DBManager) HealthCheck() error {
// 	db, err := m.GetDB()
// 	if err != nil {
// 		return err
// 	}
// 	return db.Ping()
// }

// type contextKey string

// const dbKey contextKey = "db"

// func DBMiddleware(dbManager *DBManager) func(next http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			// Lazy load database connection only when needed
// 			ctx := context.WithValue(r.Context(), dbKey, dbManager)
// 			next.ServeHTTP(w, r.WithContext(ctx))
// 		})
// 	}
// }

// func GetDBFromContext(ctx context.Context) (*DBManager, bool) {
// 	db, ok := ctx.Value(dbKey).(*DBManager)
// 	return db, ok
// }
