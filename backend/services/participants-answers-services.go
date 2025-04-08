package services

import (
	"scoremaster/backend/database"
	"scoremaster/backend/models"
)

func CreateParticipantAnswer(participantAnswer *models.ParticipantAnswer) error {
	return database.DBConn.Create(participantAnswer).Error
}
