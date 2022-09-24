package handler_test

import (
	"github.com/beto-ouverney/falemais-telzir/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"testing"
)

var allDDDCodesMock = []entity.DDDCost{
	{
		Origin:      11,
		Destination: 16,
	},
	{
		Origin:      11,
		Destination: 17,
	},
	{
		Origin:      11,
		Destination: 18,
	},
	{
		Origin:      16,
		Destination: 11,
	},
	{
		Origin:      17,
		Destination: 11,
	},
	{
		Origin:      18,
		Destination: 11,
	},
}

var schemasInit = [7]string{
	`CREATE TABLE IF NOT EXISTS dddcost(
    id SERIAL PRIMARY KEY,
    origin INT,
    destination INT,
    cost DECIMAL(10,2)
    );`,
	`INSERT INTO dddcost (origin, destination, cost) VALUES (11, 16, 1.90);`,
	`INSERT INTO dddcost (origin, destination, cost) VALUES (16, 11, 2.90);`,
	`INSERT INTO dddcost (origin, destination, cost) VALUES (11, 17, 1.70);`,
	`INSERT INTO dddcost (origin, destination, cost) VALUES (17, 11, 2.70);`,
	`INSERT INTO dddcost (origin, destination, cost) VALUES (11, 18, 0.90);`,
	`INSERT INTO dddcost (origin, destination, cost) VALUES (18, 11, 1.90);`,
}

const POSTGREES_CONNECTION = "user=root password=password dbname=telzir_db_test sslmode=disable"

func initDBTest(t *testing.T) {
	t.Log("Initializing database test")

	conn, err := sqlx.Open("postgres", POSTGREES_CONNECTION)

	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = conn.Exec("DROP TABLE IF EXISTS dddcost")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Database test dropped successfully")

	t.Setenv("POSTGRES_DB", "telzir_db_test")

	for _, s := range schemasInit {
		_, err = conn.Exec(s)
		if err != nil {
			t.Fatal(err)
		}
	}
	conn.Close()
	t.Log("Database test created successfully")
}
