package database

import (
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/lib/pq"

	migrate "github.com/rubenv/sql-migrate"
)

var (
	// host     string = os.Getenv("DB_HOST")
	// port, _         = strconv.Atoi(os.Getenv("DB_PORT"))
	// user            = os.Getenv("DB_USER")
	// password        = os.Getenv("DB_PASSWORD")
	// dbname          = os.Getenv("DB_NAME")

	host     string = "postgres.railway.internal"
	port            = 5432
	user            = "postgres"
	password        = "KlnTYuvamHUJDYQTJcGMHvsgOmNAkVAs"
	dbname          = "railway"
)

var (
	DB  *sql.DB
	err error
)

func ConnectDB() {

	pSqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err = sql.Open("postgres", pSqlInfo)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	DBMigrate(DB)

	fmt.Println("Successfully connected database!")
}

//go:embed sql_migrations/*.sql
var dbMigrations embed.FS

var DbConnection *sql.DB

func DBMigrate(dbParam *sql.DB) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migrations",
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	DbConnection = dbParam

	fmt.Println("Migration success, applied", n, "migrations!")
}
