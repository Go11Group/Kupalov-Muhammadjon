package main

import (
	"database/sql"
	"fmt"

	pgadapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize PostgreSQL adapter
	db, _ := sql.Open("postgres", "dbname=casbin user=postgres password=root host=localhost sslmode=disable")

	a, err := pgadapter.NewAdapter(db, "postgres", "casbin")
	if err != nil {
		panic(err)
	}

	// Initialize Enforcer with the adapter
	e, err := casbin.NewEnforcer("config/model.conf", a)
	if err != nil {
		panic(err)
	}

	// Load policies from the database
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}

	// Define test cases
	testCases := []struct {
		sub      string
		obj      string
		act      string
		expected bool
	}{
		{"alice", "data1", "read", true},
		{"bob", "data2", "write", false},
	}

	// Check permissions
	for _, tc := range testCases {
		result, err := e.Enforce(tc.sub, tc.obj, tc.act)
		if err != nil {
			fmt.Printf("Error enforcing policy: %v\n", err)
		} else {
			fmt.Printf("%s, %s, %s: %v\n", tc.sub, tc.obj, tc.act, result)
		}
	}

	// Add a new policy rule
	e.AddPolicy("bob", "data2", "write")
	e.SavePolicy()

	// Verify the new policy rule
	result, _ := e.Enforce("bob", "data2", "write")
	fmt.Printf("bob, data2, write: %v\n", result)
}
