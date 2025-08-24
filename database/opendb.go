package database

import "opendavinci/queries"

// Queries struct for collect all app queries.
type Queries struct {
	*queries.CourseQueries // load queries from Course model
}

// OpenDBConnection is our step to switch between diff database (pg/sqlite).
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		CourseQueries: &queries.CourseQueries{DB: db}, // from Course model
	}, nil
}

/**************************************
func OpenDBConnection111() (*Queries, error) {
	// Define a new SQLite connection.
	db, err := SqliteConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		CourseQueries: &queries.CourseQueries{DB: db}, // from Course model
	}, nil
}
**************************************/
