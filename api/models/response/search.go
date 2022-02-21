package response

import "pledge-backend/api/models"

type Search struct {
	Count int64         `json:"count"`
	Rows  []models.Pool `json:"rows"`
}
