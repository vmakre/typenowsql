package resource

import (
	"app_chi_templ/dbmanager"
	"app_chi_templ/middleware"
	"app_chi_templ/models"
	"context"
)

type MainDBResource interface {
	CreateQuery(ctx context.Context, query *models.TNQStoredQuery) error
	GetQueryByID(ctx context.Context, id int) (*models.TNQStoredQuery, error)
	GetQueries(ctx context.Context) ([]*models.TNQStoredQuery, error)
	UpdateQuery(ctx context.Context, query *models.TNQStoredQuery) error
	DeleteQuery(ctx context.Context, id int) error
}

type mainDBResource struct {
	dbManager *dbmanager.DatabaseManager
}

func NewMainDBResource(dbManager *dbmanager.DatabaseManager) MainDBResource {
	return &mainDBResource{dbManager: dbManager}
}

func (r *mainDBResource) CreateQuery(ctx context.Context, query *models.TNQStoredQuery) error {
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
	queryq := `INSERT INTO TNQ_StoredQueries ([QueryName] ,[QueryText],[Description],[LastModified],[Parameters],[IsActive])
     	VALUES ($1, $2 , $3 ,$4 , $5 , $6 ) `
	return db.QueryRowContext(ctx, queryq, query.QueryName, query.QueryText).Scan()
	// return r.db.QueryRowContext(ctx, query, user.Name, user.Email).Scan(
	// 	&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *mainDBResource) GetQueryByID(ctx context.Context, id int) (*models.TNQStoredQuery, error) {
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
	queryq := `select QueryID , QueryName, QueryText, Description, LastModified, Parameters, IsActive  from TNQ_StoredQueries WHERE QueryID = $1`
	row := db.QueryRowContext(ctx, queryq, id)

	var query models.TNQStoredQuery
	err = row.Scan(&query.QueryID, &query.QueryName, &query.QueryText, &query.Description, &query.Parameters)
	if err != nil {
		return nil, err
	}
	return &query, nil
}

func (r *mainDBResource) GetQueries(ctx context.Context) ([]*models.TNQStoredQuery, error) {
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
	queryq := `select QueryID , QueryName, QueryText, Description, LastModified, Parameters, IsActive  from TNQ_StoredQueries`
	rows, err := db.QueryContext(ctx, queryq)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var querries []*models.TNQStoredQuery
	for rows.Next() {
		var query models.TNQStoredQuery
		if err := rows.Scan(&query.QueryID, &query.QueryName, &query.QueryText, &query.Description, &query.Parameters); err != nil {
			return nil, err
		}
		querries = append(querries, &query)
	}
	return querries, nil
}

// func (r *mainDBResource) GetUsers(ctx context.Context) ([]*models.User, error) {
// 	dbManager := middleware.GetDBManagerFromContext(ctx.Value("context").(context.Context))
// 	if dbManager == nil {
// 		//http.Error(w, "Database manager not available", http.StatusInternalServerError)
// 		return nil, nil
// 	}

// 	// Use specific database connection
// 	db, err := dbManager.GetConnection("main")
// 	if err != nil {
// 		// http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return nil, nil
// 	}
// 	//query := `SELECT id, name, email, created_at, updated_at FROM users`
// 	query := `SELECT city , country_id  FROM city`
// 	rows, err := db.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []*models.User
// 	for rows.Next() {
// 		var user models.User
// 		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
// 			return nil, err
// 		}
// 		users = append(users, &user)
// 	}
// 	return users, nil
// }

func (r *mainDBResource) UpdateQuery(ctx context.Context, query *models.TNQStoredQuery) error {
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
	queryq := `UPDATE TNQ_StoredQueries SET QueryName = $1, QueryText = $2, QueryID = $3, updated_at = NOW() WHERE id = $4`
	_, err = db.ExecContext(ctx, queryq, query.QueryName, query.QueryText, query.Parameters, query.QueryID)
	return err
}

func (r *mainDBResource) DeleteQuery(ctx context.Context, id int) error {
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
	query := `DELETE FROM TNQ_StoredQueries WHERE id = $1`
	_, err = db.ExecContext(ctx, query, id)
	return err
}
