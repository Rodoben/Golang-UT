package dbrepo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
	"web-testing/pkg/data"
	"web-testing/pkg/repository"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbName   = "users_test"
	port     = "5435"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var (
	resource *dockertest.Resource
	pool     *dockertest.Pool
	testDB   *sql.DB
	testRepo repository.DatabaseRepo
)

func TestMain(m *testing.M) {
	// connect to docker; fail id not running
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker, is it running?")
	}
	pool = p
	// setup our docker options, specifying the image and do fourth
	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.5",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbName,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}
	// get a resource (docker image)
	resource, err = pool.RunWithOptions(&opts)
	if err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not start resource: %s", err)
	}
	// start the image and wait until its ready
	if err := pool.Retry(func() error {
		var err error
		testDB, err = sql.Open("pgx", fmt.Sprintf(dsn, host, port, user, password, dbName))
		if err != nil {
			log.Println("Error:", err)
			return err
		}
		return testDB.Ping()
	}); err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("Could not connect to database: %s", err)
	}

	// populate the database with empty tables
	err = createTables()
	if err != nil {
		log.Fatalf("Error Creating tables: %s", err)

	}
	testRepo = &PostgresDBRepo{DB: testDB}
	//cleanup

	code := m.Run()
	if resource != nil {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("could not purge resources: %s", err)
		}
	}
	os.Exit(code)
}

func createTables() error {
	tableSQL, err := os.ReadFile("./testdata/users.sql")
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	_, err = testDB.Exec(string(tableSQL))
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	fmt.Println("Tables created")
	return nil
}

func Test_pingDB(t *testing.T) {
	err := testDB.Ping()
	if err != nil {
		t.Error("Cant ping database")
	}
}

func TestPostgresDBRepoInsertUser(t *testing.T) {
	testUser := data.User{
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@example.com",
		Password:  "secret",
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	id, err := testRepo.InsertUser(testUser)
	if err != nil {
		t.Errorf("insert user returned an error: %s", err)
	}

	if id != 1 {
		t.Errorf("insert user returned wrong id; expected 1, but got %d", id)
	}
}
func TestPostgresDBRepo_Table_InsertUser(t *testing.T) {
	insertTest := []struct {
		name     string
		input    data.User
		response int
	}{
		{name: "success", input: data.User{
			ID:        2,
			FirstName: "Admin",
			LastName:  "User",
			Email:     "admin@example.com",
			Password:  "secret",
			IsAdmin:   1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
			response: 2,
		},
		{name: "success", input: data.User{
			ID:        3,
			FirstName: "Admin",
			LastName:  "User",
			Email:     "admin@example.com",
			Password:  "secret",
			IsAdmin:   1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
			response: 3,
		},
	}

	for _, v := range insertTest {
		id, err := testRepo.InsertUser(v.input)
		if err != nil {
			t.Errorf("insert user returned error %s:", err)
		}
		if id != v.response {
			t.Errorf("insert user returned wrong id; expected %d: but got %d", v.response, id)
		}
	}
}

func TestPostgresDBRepo_Table_ALLUser(t *testing.T) {
	insertTest := []struct {
		name string
	}{
		{name: "success"},
	}

	for range insertTest {
		users, err := testRepo.AllUsers()
		if err != nil {
			t.Errorf("All user returned error %s:", err)
		}
		if len(users) == 0 {
			t.Errorf("All user returned invalid length; expected %d: but got %d", 2, len(users))
		}
	}
}

func TestPostgresDBRepo_Table_GetUser(t *testing.T) {
	testGetUser := []struct {
		name     string
		input    int
		response string
	}{
		{name: "valid id", input: 1, response: "admin@example.com"},
	}

	for _, v := range testGetUser {
		user, err := testRepo.GetUser(v.input)
		if err != nil {
			t.Errorf("Get user returned error: %s", err)
		}

		if user.ID != v.input {
			t.Errorf("GetUser returned invalid length; expected %d: but got %d", v.input, user.ID)
		}
		if user.Email != v.response {
			t.Errorf("GetUser returned invalid email; expected %s: but got %s:", v.response, user.Email)
		}

	}

	_, err := testRepo.GetUser(5)
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}
}

func TestPostgresDBRepo_Table_GetUserByEmail(t *testing.T) {
	testGetUserByEmail := []struct {
		name     string
		input    string
		response string
	}{
		{name: "valid email", input: "admin@example.com", response: "admin@example.com"},
	}

	for _, v := range testGetUserByEmail {
		user, err := testRepo.GetUserByEmail(v.input)
		if err != nil {
			t.Errorf("Get user returned error: %s", err)
		}

		if user.Email != v.response {
			t.Errorf("GetUser returned invalid email; expected %s: but got %s:", v.response, user.Email)
		}

	}

}
