package pet

import (
	"encoding/json"

	"github.com/wexel-nath/pets/pkg/database"
	"github.com/wexel-nath/pets/pkg/logger"
)

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Pet struct {
	ID        int64                    `json:"id"`
	Category  Category                 `json:"category"`
	Name      string                   `json:"name"`
	PhotoURLs []map[string]interface{} `json:"photoUrls"`
	Tags      []Tag                    `json:"tags"`
	Status    string                   `json:"status"`
}

func NewPet(row map[string]interface{}) (Pet, error) {
	pet := Pet{
		ID: row["pet_id"].(int64),
		Category: Category{
			ID:   row["pet_category_id"].(int64),
			Name: row["pet_category_name"].(string),
		},
		Name:   row["pet_name"].(string),
		Status: row["pet_status_name"].(string),
	}

	if err := json.Unmarshal(row["pet_photo_urls"].([]byte), &pet.PhotoURLs); err != nil {
		return Pet{}, err
	}
	if err := json.Unmarshal(row["pet_tags"].([]byte), &pet.Tags); err != nil {
		return Pet{}, err
	}

	return pet, nil
}

func Insert(pet Pet) (Pet, error) {
	db := database.GetConnection()
	tx, err := db.Begin()
	if err != nil {
		return Pet{}, err
	}

	row, err := insert(tx, pet)
	if err != nil {
		logger.LogIfErr(tx.Rollback())
		return Pet{}, err
	}

	if err = tx.Commit(); err != nil {
		return Pet{}, err
	}

	return NewPet(row)
}
