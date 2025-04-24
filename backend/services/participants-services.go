package services

import (
	"errors"
	"scoremaster/backend/database"
	"scoremaster/backend/models"
)

func CreateParticipant(participant *models.Participant) error {
	var existing models.Participant
	result := database.DBConn.Where("email=?", participant.Email).First(&existing)

	if result.RowsAffected > 0 {
        return errors.New("participant with this email already exists")
	}
	return database.DBConn.Create(participant).Error
}

// CheckEmailExists verifies if an email is already registered
func CheckEmailExists(email string) (bool, error) {
    var existing models.Participant
    result := database.DBConn.Where("email = ?", email).First(&existing)
    
    return result.RowsAffected > 0, result.Error
}

