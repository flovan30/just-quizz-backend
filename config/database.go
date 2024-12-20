package config

import (
	"just-quizz-server/database"
)

func InitDatabaseConnection() {
	// initialize database connection

	database.InitDB()
}
