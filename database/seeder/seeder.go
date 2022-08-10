package seeder

import (
	"github.com/agusbasari29/simple-xjx-tasking.git/database"
	"github.com/agusbasari29/simple-xjx-tasking.git/repository"
)

var (
	db       = database.SetupDatabaseConnection()
	taskRepo = repository.NewTasksRepository(db)
)

func LetsSeed() {
	TasksSeederUp(10)
}
