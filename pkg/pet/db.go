package pet

import (
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/wexel-nath/pets/pkg/database"
	"github.com/wexel-nath/pets/pkg/logger"
)

var (
	selectColumns = []string{
		"pet_id",
		"pet_name",
		"pet_category_id",
		"pet_category_name",
		"pet_photo_urls",
		"pet_tags",
		"pet_status_name",
	}

	insertColumns = []string{
		"pet_name",
		"pet_category_id",
		"pet_photo_urls",
		"pet_tags",
		"pet_status_id",
	}

	petStatusMap = map[string]int64{
		"available": 10,
		"pending":   20,
		"sold":      30,
	}
)

func insert(tx *sql.Tx, pet Pet) (map[string]interface{}, error) {
	upsertCategoryQuery := `
		WITH spc AS (
			SELECT pet_category_id
			FROM   pet_category
			WHERE  pet_category_name = $1
		),
		ipc AS (
			INSERT INTO pet_category (pet_category_name)
			SELECT $1
			WHERE NOT EXISTS (SELECT 1 FROM spc)
			RETURNING pet_category_id
		)
		SELECT pet_category_id
		FROM ipc
		UNION ALL
		SELECT pet_category_id
		FROM spc
	`

	var categoryID int64
	err := tx.QueryRow(upsertCategoryQuery, pet.Category.Name).Scan(&categoryID)
	if err != nil {
		return nil, err
	}

	insertPetQuery := `
		WITH select_pet AS (
			INSERT INTO pet (
				` + strings.Join(insertColumns, ", ") + `
			)
			VALUES (
				$1, $2, $3, $4, $5
			)
			RETURNING *
		)
		SELECT
			` + strings.Join(selectColumns, ", ") + `
		FROM select_pet
			JOIN pet_category USING (pet_category_id)
			JOIN pet_status USING (pet_status_id)
	`

	photoUrls, err := json.Marshal(pet.PhotoURLs)
	if err != nil {
		logger.LogIfErr(tx.Rollback())
		return nil, err
	}

	tags, err := json.Marshal(pet.Tags)
	if err != nil {
		return nil, err
	}

	params := []interface{}{
		pet.Name,
		categoryID,
		photoUrls,
		tags,
		petStatusMap[pet.Status],
	}

	row := tx.QueryRow(insertPetQuery, params...)
	return database.ScanRowToMap(row, selectColumns)
}
