package models

import "database/sql"

type TNQStoredQuery struct {
	QueryID      int64          `json:"queryId" db:"query_id"`
	QueryName    string         `json:"queryName" db:"query_name"`
	QueryText    string         `json:"queryText" db:"query_text"`
	Description  sql.NullString `json:"description" db:"description"`
	LastModified sql.NullTime   `json:"lastModified" db:"last_modified"`
	Parameters   sql.NullString `json:"parameters" db:"parameters"`
	IsActive     sql.NullBool   `json:"isActive" db:"is_active"`
}
