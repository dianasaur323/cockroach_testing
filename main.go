import(
  "database/sql"
  "fmt"
)

var concurrency = flag.Int("concurrency", 2*runtime.NumCPU(), "Number of concurrent writers inserting blocks")

func setupDatabase(dbURL string) (*sql.DB, error){
  fmt.Printf("connecting to db: %s\n", dbURL)
  parsedURL, err := url.Parse(dbURL)

  // Open connection to server and create a database.
  db, err  := sql.Open("postgres", parsedURL.String())
  if err != nil {
    return nil, err
  }

  // Allow a maximum of concurrency connections to database.
  db.SetMaxOpenConns(*concurrency)
  db.SetMaxIdleConns(*concurrency)

  return db, nil
}
