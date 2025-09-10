CREATE TABLE TNQ_StoredQueries(
    QueryID INTEGER PRIMARY KEY AUTOINCREMENT,
    QueryName TEXT NOT NULL,
    QueryText TEXT NOT NULL,
    Description TEXT NULL,
    LastModified DATETIME NULL DEFAULT (datetime('now')),
    Parameters TEXT NULL,
    IsActive INTEGER NULL DEFAULT 1
);