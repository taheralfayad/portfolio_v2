package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
	utils "github.com/taheralfayad/portfolio_v2/utils"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "create":
		createUser()
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: go run create_user.go <command>")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  create    Create a new user")
	fmt.Println("  help      Show this help message")
}

func createUser() {
	username := os.Getenv("USER_NAME")
	password := os.Getenv("USER_PASSWORD")

	// If environment variables are not set, fall back to interactive input
	if username == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter username: ")
		usernameInput, _ := reader.ReadString('\n')
		username = strings.TrimSpace(usernameInput)
	}

	if password == "" {
		fmt.Print("Enter password: ")
		passwordInput, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			// Fallback to regular input if term.ReadPassword fails
			fmt.Println()
			fmt.Print("Enter password (visible): ")
			reader := bufio.NewReader(os.Stdin)
			passwordInput, _ := reader.ReadString('\n')
			password = strings.TrimSpace(string(passwordInput))
		} else {
			fmt.Println()
			password = string(passwordInput)
		}
	}

	if username == "" {
		log.Fatal("Username cannot be empty")
	}

	if password == "" {
		log.Fatal("Password cannot be empty")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	db, err := connectToDatabase()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	var userID int
	var createdAt string

	query := `
		INSERT INTO users (name, password)
		VALUES ($1, $2)
		RETURNING id, created_at
	`

	err = db.QueryRow(query, username, hashedPassword).Scan(&userID, &createdAt)
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	fmt.Printf("User created successfully!\n")
	fmt.Printf("ID: %d\n", userID)
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Created At: %s\n", createdAt)
}

func connectToDatabase() (*sql.DB, error) {
	dbUser := getEnv("DB_USER", "taher")
	dbPassword := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "portfolio")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
