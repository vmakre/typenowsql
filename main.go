package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"typenowsql/dbmanager"
	"typenowsql/handlers"
	"typenowsql/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	. "maragu.dev/gomponents"
	ghttp "maragu.dev/gomponents/http"

	// "github.com/gohugoio/hugo/resources/resource"
	"maragu.dev/httph"
)

func main() {
	// Database configuration
	// dbConfig := database.DBConfig{
	// 	Host:     getEnv("DB_HOST", "10.10.4.160"),
	// 	Port:     getEnvAsInt("DB_PORT", 1435),
	// 	User:     getEnv("DB_USER", "sa"),
	// 	Password: getEnv("DB_PASSWORD", "M@rek2017"),
	// 	DBName:   getEnv("DB_NAME", "sakila"),
	// 	SSLMode:  getEnv("DB_SSLMODE", "disable"),
	// }
	// Create database manager with lazy loading
	//	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	//dbManager := database.NewDBManager(dsn, 25, 5, time.Hour, false)

	// Initialize database manager
	dbManager := dbmanager.NewDatabaseManager()
	dbMain := dbmanager.DBConfig{
		Type:   dbmanager.SQLite,
		DBName: "sqlite.db",
	}
	if err := dbManager.AddConnection("main", dbMain); err != nil {
		//log.Fatal(err)
	}
	// Connect to databases
	if err := dbManager.Connect("main"); err != nil {
		//	log.Fatal(err)
	}

	// Initialize handlers
	dbHandler := handlers.NewDBHandler(dbManager)

	// Create Chi router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(httph.VersionedAssets)
	// Get DB connection from manager and set up middleware
	//r.Use(database.DBMiddleware(dbManager)) // Inject DB into context
	Static(r)

	// Routes
	r.Get("/health", dbHandler.HealthCheckHandler)
	r.Get("/query/{dbName}", dbHandler.QueryHandler)

	// Initialize repository, service, and handlers with dependency injection
	// userRepo := resource.NewUserResource(dbHandler.QueryHandler)
	// userService := service.NewUserService(userRepo)
	// userHandler := handlers.NewUserHandler(userService)
	homepageService := service.NewHomePageService()
	homePageHandler := handlers.NewHomePageHandler(homepageService) // Pass actual service if needed

	// Routes
	// r.Route("/api/users", func(r chi.Router) {
	// 	r.Post("/", userHandler.CreateUser)
	// 	r.Get("/", userHandler.GetUsers)
	// 	r.Get("/{id}", userHandler.GetUser)
	// 	r.Put("/{id}", userHandler.UpdateUser)
	// 	r.Delete("/{id}", userHandler.DeleteUser)
	// })
	r.Route("/query", func(r chi.Router) {
		r.Get("/", homePageHandler.GetGueryPage())
		r.Post("/{dbName}", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
			return dbHandler.QueryHandlerHtmlTable(w, r)
		}))
	})

	r.Get("/", homePageHandler.GetHomePage())

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Start server
	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// Helper functions for environment variables
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// Static assets handler, which serves files from the root that have an extension, and everything from
// the images, scripts, and styles directories.
func Static(r chi.Router) {
	staticHandler := http.FileServer(http.Dir("public"))
	r.Get(`/{:[^.]+\.[^.]+}`, staticHandler.ServeHTTP)
	r.Get(`/{:images|scripts|styles}/*`, staticHandler.ServeHTTP)
}
