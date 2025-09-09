package handlers

import (
	"app_chi_templ/dbmanager"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type DBHandler struct {
	dbManager *dbmanager.DatabaseManager
}

func NewDBHandler(dbManager *dbmanager.DatabaseManager) *DBHandler {
	return &DBHandler{dbManager: dbManager}
}

// HealthCheckHandler returns the health status of all connections
func (h *DBHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	health := h.dbManager.HealthCheck()

	response := make(map[string]string)
	for name, err := range health {
		if err != nil {
			response[name] = "unhealthy: " + err.Error()
		} else {
			response[name] = "healthy"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// QueryHandler executes a query on a specific database
func (h *DBHandler) QueryHandler(w http.ResponseWriter, r *http.Request) {
	dbName := chi.URLParam(r, "dbName")
	query := r.URL.Query().Get("query")

	if query == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	db, err := h.dbManager.GetConnection(dbName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	ctx := r.Context()
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var results []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := make(map[string]interface{})
		for i, col := range columns {
			result[col] = values[i]
		}
		results = append(results, result)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// QueryHandler executes a query on a specific database
func (h *DBHandler) QueryHandlerHtmlTable(w http.ResponseWriter, r *http.Request) (Node, error) {
	dbName := chi.URLParam(r, "dbName")
	//query := r.URL.Query().Get("query")
	// json decode
	var data struct {
		Query string `json:"query"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return nil, err
	}

	db, err := h.dbManager.GetConnection(dbName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return nil, nil
	}
	resultSetIndex := 0
	ctx := r.Context()
	rows, err := db.QueryContext(ctx, data.Query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	defer rows.Close()
	var columnsx [][]string
	var resultSet [][]map[string]interface{}
	for { // Loop for each result set
		resultSetIndex++
		fmt.Printf("--- Processing Result Set #%d ---\n", resultSetIndex)
		columns, err := rows.Columns()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil, err
		}
		var results []map[string]interface{}
		for rows.Next() {
			values := make([]interface{}, len(columns))
			valuePtrs := make([]interface{}, len(columns))
			for i := range values {
				valuePtrs[i] = &values[i]
			}
			if err := rows.Scan(valuePtrs...); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return nil, err
			}

			result := make(map[string]interface{})
			for i, col := range columns {
				result[col] = values[i]
			}
			results = append(results, result)
		}
		columnsx = append(columnsx, columns)
		resultSet = append(resultSet, results)
		// Try to move to the next result set. Break loop if none.
		if !rows.NextResultSet() {
			break
		}

	}
	var increment int = -1
	return Group{
		Map(resultSet, func(r []map[string]interface{}) Node {
			increment++
			if len(r) == 0 {
				return Div(Text("No results"))
			}
			columns := columnsx[increment]
			results := r
			return Div(
				Span(P(Style("background-color:#842626;padding:2px;"))),
				Class("overflow-x-auto"),
				Style("margin-bottom:10px"),
				Table(Class("table table-xs"),
					THead(
						Tr(
							Group(
								Map(columns, func(t string) Node {
									return Th(Text(t))
								}),
							),
						),
					),
					TBody(
						Group(
							Map(results, func(row map[string]interface{}) Node {
								return Tr(Class("bg-base-300"),
									Group(
										Map(columns, func(col string) Node {
											value := row[col]
											return Td(Text(fmt.Sprintf("%v", value)))
										}),
									),
								)
							}),
						),
					),
				),
			)
		}),
	}, nil

}

// QueryHandler executes a query on a specific database
func (h *DBHandler) DBConnsHandlerHtmlTable(w http.ResponseWriter, r *http.Request) (Node, error) {
	// dbName := chi.URLParam(r, "dbName")

	return nil, nil
	// return Group{
	// 	Div(
	// 		Select(Class("select select-bordered w-full max-w-xs"),
	// 			Map(h.dbManager.GetConnections(), func(r dbmanager.DatabaseManager) Node {
	// 				if r.Name == dbName {
	// 					return Option(Attr("value", r.Name), Text(r.Name))
	// 				}
	// 				return nil
	// 			},
	// 			)),
	// 	),
	// }, nil
}
