package services

import (
	"scoremaster/backend/database"
	"scoremaster/backend/models"
)

func CreateParticipant(participant *models.Participant) error {
	return database.DBConn.Create(participant).Error
}
