package main

import (
	"time"

	"github.com/Paienobe/go-todo/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Password  string    `json:"password"`
}

func dbUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		Email:     dbUser.Email,
		CreatedAt: dbUser.CreatedAt,
		Password:  dbUser.Password,
	}
}

type Task struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Iscompleted bool      `json:"isCompleted"`
	UserID      uuid.UUID `json:"userId"`
}

func dbTaskToTask(dbTask database.Task) Task {
	return Task{
		ID:          dbTask.ID,
		Name:        dbTask.Name,
		Iscompleted: dbTask.Iscompleted,
		UserID:      dbTask.UserID,
	}
}

func dbTasksToTasks(dbTasks []database.Task) []Task {
	tasks := []Task{}
	for _, dbTask := range dbTasks {
		tasks = append(tasks, Task{
			ID:          dbTask.ID,
			Name:        dbTask.Name,
			Iscompleted: dbTask.Iscompleted,
			UserID:      dbTask.UserID,
		})
	}
	return tasks
}
