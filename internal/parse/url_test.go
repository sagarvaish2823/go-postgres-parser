package parse

import "testing"

func TestPostgresURL(t *testing.T) {
	f, err := PostgresURL("postgres://neondb_owner:secret@ep-cool.us-east-2.aws.neon.tech:5432/neondb?sslmode=require")
	if err != nil {
		t.Fatal(err)
	}
	if f.Host != "ep-cool.us-east-2.aws.neon.tech" {
		t.Fatalf("host: %q", f.Host)
	}
	if f.Port != "5432" {
		t.Fatalf("port: %q", f.Port)
	}
	if f.Database != "neondb" {
		t.Fatalf("database: %q", f.Database)
	}
	if f.Username != "neondb_owner" {
		t.Fatalf("username: %q", f.Username)
	}
	if f.Password != "secret" {
		t.Fatalf("password: %q", f.Password)
	}
}

func TestPostgresURLDefaultPort(t *testing.T) {
	f, err := PostgresURL("postgresql://user:pass@localhost/mydb")
	if err != nil {
		t.Fatal(err)
	}
	if f.Port != "5432" {
		t.Fatalf("port: %q", f.Port)
	}
}
