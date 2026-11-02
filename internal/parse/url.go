package parse

import (
	"fmt"
	"net/url"
	"strconv"
)

type Fields struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func PostgresURL(raw string) (Fields, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return Fields{}, fmt.Errorf("invalid URL: %w", err)
	}
	if u.Scheme != "postgres" && u.Scheme != "postgresql" {
		return Fields{}, fmt.Errorf("URL scheme must be postgres or postgresql")
	}

	host := u.Hostname()
	if host == "" {
		return Fields{}, fmt.Errorf("URL is missing host")
	}

	port := u.Port()
	if port == "" {
		port = "5432"
	} else if _, err := strconv.Atoi(port); err != nil {
		return Fields{}, fmt.Errorf("invalid port %q", port)
	}

	db := u.Path
	if len(db) > 0 && db[0] == '/' {
		db = db[1:]
	}
	if db == "" {
		return Fields{}, fmt.Errorf("URL is missing database name")
	}

	user := u.User.Username()
	if user == "" {
		return Fields{}, fmt.Errorf("URL is missing username")
	}
	pass, _ := u.User.Password()

	return Fields{
		Host:     host,
		Port:     port,
		Database: db,
		Username: user,
		Password: pass,
	}, nil
}
