package controller



import (

    "testing"

)



func TestDBConnection(t *testing.T) {

    if db == nil {

        t.Fatal("db is nil")

    }

    if err := db.Ping(); err != nil {

        t.Fatalf("db ping failed: %v", err)

    }

    // close DB after test to avoid leaking connections

    if err := CloseDB(); err != nil {

        t.Fatalf("failed to close db: %v", err)

    }

}