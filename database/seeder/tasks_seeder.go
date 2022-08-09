package seeder

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/agusbasari29/simple-xjx-tasking.git/entity"
	"github.com/bxcodec/faker"
)

type TasksSeeder struct {
	Task     string `faker:"sentence"`
	Assignee string `faker:"name"`
	Deadline int64  `faker:"unix_time"`
}

func TasksSeederUp(number int) {
	seeder := TasksSeeder{}
	tasks := entity.Tasks{}
	for i := 0; i < number; i++ {
		var TaskStatus entity.TaskStatus
		j := rand.Intn(2)
		err := faker.FakeData(&seeder)
		if err != nil {
			fmt.Printf("%+v", err)
		}
		switch j {
		case 1:
			TaskStatus = entity.Completed
		case 2:
			TaskStatus = entity.Idle
		default:
			TaskStatus = entity.Progress
		}
		tasks.Status = TaskStatus
		tasks.Task = seeder.Task
		tasks.Assignee = seeder.Assignee
		tasks.Deadline = convertTime(seeder.Deadline)
		taskRepo.CreateTask(tasks)
	}
}

func convertTime(unix int64) time.Time {
	i, err := strconv.ParseInt(strconv.Itoa(int(unix)), 10, 64)
	if err != nil {
		panic(err)
	}
	return time.Unix(i, 0)
}
