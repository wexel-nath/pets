package user

import (
	"strings"

	"github.com/wexel-nath/pets/pkg/database"
)

var (
	selectColumns = []string{
		"user_id",
		"username",
		"first_name",
		"last_name",
		"email",
		"password",
		"phone",
		"status",
	}

	insertColumns = []string{
		"username",
		"first_name",
		"last_name",
		"email",
		"password",
		"phone",
	}
)

func insert(user User) (map[string]interface{}, error) {
	query := `
		INSERT INTO "user" (
			` + strings.Join(insertColumns, ", ") + `
		)
		VALUES (
			$1, $2, $3, $4, $5, $6
		)
		RETURNING *
	`

	params := []interface{}{
		user.Username,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Phone,
	}

	db := database.GetConnection()
	row := db.QueryRow(query, params...)
	return database.ScanRowToMap(row, selectColumns)
}
