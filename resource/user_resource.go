package resource

import (
	"app_chi_templ/dbmanager"
	"app_chi_templ/middleware"
	"app_chi_templ/models"
	"context"
)

type UserResource interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	GetCities(ctx context.Context) ([]*models.City, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int) error
}

// type DBHandler struct {
// 	dbManager *dbmanager.DatabaseManager
// }

// func NewDBHandler(dbManager *dbmanager.DatabaseManager) *DBHandler {
// 	return &DBHandler{dbManager: dbManager}
// }

type userResource struct {
	// db *sql.DB
	dbManager *dbmanager.DatabaseManager
}

func NewUserResource(dbManager *dbmanager.DatabaseManager) UserResource {
	return &userResource{dbManager: dbManager}
}

func (r *userResource) CreateUser(ctx context.Context, user *models.User) error {
	dbManager := middleware.GetDBManagerFromContext(ctx.Value("context").(context.Context))
	if dbManager == nil {
		//http.Error(w, "Database manager not available", http.StatusInternalServerError)
		return nil
	}

	// Use specific database connection
	db, err := dbManager.GetConnection("main")
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at, updated_at`
	return db.QueryRowContext(ctx, query, user.Name, user.Email).Scan(
		&user.ID, &user.CreatedAt, &user.UpdatedAt)

	// return r.db.QueryRowContext(ctx, query, user.Name, user.Email).Scan(
	// 	&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *userResource) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	dbManager := middleware.GetDBManagerFromContext(ctx.Value("context").(context.Context))
	if dbManager == nil {
		//http.Error(w, "Database manager not available", http.StatusInternalServerError)
		return nil, nil
	}

	// Use specific database connection
	db, err := dbManager.GetConnection("main")
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, nil
	}
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1`
	row := db.QueryRowContext(ctx, query, id)

	var user models.User
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userResource) GetCities(ctx context.Context) ([]*models.City, error) {
	dbManager := middleware.GetDBManagerFromContext(ctx)
	if dbManager == nil {
		//http.Error(w, "Database manager not available", http.StatusInternalServerError)
		return nil, nil
	}

	// Use specific database connection
	db, err := dbManager.GetConnection("main")
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, nil
	}
	//query := `SELECT id, name, email, created_at, updated_at FROM users`
	query := `SELECT city , country_id  FROM city`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []*models.City
	for rows.Next() {
		var city models.City
		if err := rows.Scan(&city.City, &city.Country_id); err != nil {
			return nil, err
		}
		cities = append(cities, &city)
	}
	return cities, nil
}
func (r *userResource) GetUsers(ctx context.Context) ([]*models.User, error) {
	dbManager := middleware.GetDBManagerFromContext(ctx.Value("context").(context.Context))
	if dbManager == nil {
		//http.Error(w, "Database manager not available", http.StatusInternalServerError)
		return nil, nil
	}

	// Use specific database connection
	db, err := dbManager.GetConnection("main")
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, nil
	}
	//query := `SELECT id, name, email, created_at, updated_at FROM users`
	query := `SELECT city , country_id  FROM city`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *userResource) UpdateUser(ctx context.Context, user *models.User) error {
	dbManager := middleware.GetDBManagerFromContext(ctx.Value("context").(context.Context))
	if dbManager == nil {
		//http.Error(w, "Database manager not available", http.StatusInternalServerError)
		return nil
	}

	// Use specific database connection
	db, err := dbManager.GetConnection("main")
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	query := `UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3`
	_, err = db.ExecContext(ctx, query, user.Name, user.Email, user.ID)
	return err
}

func (r *userResource) DeleteUser(ctx context.Context, id int) error {
	dbManager := middleware.GetDBManagerFromContext(ctx.Value("context").(context.Context))
	if dbManager == nil {
		//http.Error(w, "Database manager not available", http.StatusInternalServerError)
		return nil
	}

	// Use specific database connection
	db, err := dbManager.GetConnection("main")
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	query := `DELETE FROM users WHERE id = $1`
	_, err = db.ExecContext(ctx, query, id)
	return err
}
