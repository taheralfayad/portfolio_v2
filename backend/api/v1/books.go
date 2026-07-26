package v1

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	_ "github.com/lib/pq"

	data "github.com/taheralfayad/portfolio_v2/data"
	"github.com/taheralfayad/portfolio_v2/messages"
)

func GetBooks(c *gin.Context, db *sql.DB) {
	var books []data.BookMeta

	query := `
		SELECT 
			title,
			author,
			status,
			percentage_finished
		FROM book
	`

	rows, err := db.Query(query)

	if err != nil {
		slog.Error("error while retrieving rows", err)
		messages.InternalError(c, errors.New("sorry, an error occurred while retrieving these rows"))
	}

	for rows.Next() {
		var book data.BookMeta

		rows.Scan(&book.Title,
			&book.Authors,
			&book.Status,
			&book.PercentFinished,
		)

		books = append(books, book)
	}

	messages.StatusOk(c, books)
}

func PostBooks(c *gin.Context, db *sql.DB) {
	var payload data.PostBooksPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		slog.Error("error while receiving books", err)
		return
	}

	currentlyReadingFilename := strings.TrimSuffix(payload.CurrentlyReading.Text, ".epub")
	lastReadTime := payload.CurrentlyReading.MandatoryTime

	for filename, bookMeta := range payload.Library {
		if filename == currentlyReadingFilename {
			bookMeta.Status = "currently_reading"
			bookMeta.LastReadTime = lastReadTime
			payload.Library[filename] = bookMeta
		}
	}

	titles := make([]string, 0, len(payload.Library))
	authors := make([]string, 0, len(payload.Library))

	for _, bookMeta := range payload.Library {
		query := `
			INSERT INTO book 
			(title, author, last_read_time, percentage_finished, status)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (title, author) DO UPDATE SET
				author = EXCLUDED.author,
				last_read_time = EXCLUDED.last_read_time,
				percentage_finished = EXCLUDED.percentage_finished,
				status = EXCLUDED.status
		`
		if bookMeta.PercentFinished < 0.03 || bookMeta.Status == "" {
			bookMeta.Status = "not_yet_read"
		}
		_, err := db.Exec(
			query,
			bookMeta.Title,
			bookMeta.Authors,
			bookMeta.LastReadTime,
			bookMeta.PercentFinished,
			bookMeta.Status,
		)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		titles = append(titles, bookMeta.Title)
		authors = append(authors, bookMeta.Authors)
	}

	deleteQuery := `
		DELETE FROM book
		WHERE NOT (title, author) IN (
			SELECT * FROM unnest($1::text[], $2::text[])
		)
	`
	_, err := db.Exec(deleteQuery, pq.Array(titles), pq.Array(authors))
	if err != nil {
		fmt.Println("Error deleting stale books:", err)
		return
	}
}
