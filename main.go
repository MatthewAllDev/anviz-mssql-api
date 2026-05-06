package main

import (
	"anviz-mssql-api/api"
	"anviz-mssql-api/api/auth"
	"anviz-mssql-api/db"
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

func main() {
	loadDotEnv(".env")

	dsn, err := loadSecret(os.Getenv("DB_DSN"), os.Getenv("DB_DSN_FILE"), "DB_DSN_FILE or DB_DSN is required")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.OpenSQLServer(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	apiKeys, err := auth.LoadAPIKeys(os.Getenv("API_KEYS"), os.Getenv("API_KEYS_FILE"))
	if err != nil {
		log.Fatalf("Failed to load API keys: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	api.StartServer(conn, apiKeys, port)
}

func loadSecret(envValue, filename, requiredMessage string) (string, error) {
	if strings.TrimSpace(filename) != "" {
		content, err := os.ReadFile(filename)
		if err != nil {
			return "", err
		}
		value := strings.TrimSpace(string(content))
		if value == "" {
			return "", os.ErrInvalid
		}
		return value, nil
	}
	value := strings.TrimSpace(envValue)
	if value == "" {
		return "", errors.New(requiredMessage)
	}
	return value, nil
}

func loadDotEnv(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		value = strings.Trim(strings.TrimSpace(value), `"'`)
		if key == "" || os.Getenv(key) != "" {
			continue
		}
		if err := os.Setenv(key, value); err != nil {
			log.Printf("failed to load .env variable %s: %v", key, err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("failed to read .env: %v", err)
	}
}
