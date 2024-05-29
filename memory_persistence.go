package main

type MemoryPersistence struct {
	exercises []Exercise
}

func getMaxId(exercises []Exercise) int64 {
	maxId := int64(0)
	for _, exercise := range exercises {
		if exercise.ID > maxId {
			maxId = exercise.ID
		}
	}
	return maxId
}

func (p *MemoryPersistence) addExercise(exercise ExerciseCreate) bool {
	nextId := getMaxId(p.exercises)
	p.exercises = append(p.exercises, Exercise{
		ID:          nextId + 1,
		Name:        exercise.Name,
		Description: exercise.Description,
	})
	return true
}

func (p *MemoryPersistence) deleteExercise(id int64) bool {
	for i := 0; i < len(p.exercises); i++ {
		if p.exercises[i].ID == id {
			p.exercises = append(p.exercises[:i], p.exercises[i+1:]...)
			return true
		}
	}
	return false
}

func (p *MemoryPersistence) updateExerciseById(id int64, exercise ExerciseUpdate) *Exercise {
	for i := 0; i < len(p.exercises); i++ {
		if p.exercises[i].ID == id {
			if exercise.Name != "" {
				p.exercises[i].Name = exercise.Name
			}
			if exercise.Description != "" {
				p.exercises[i].Description = exercise.Description
			}

			return &p.exercises[i]
		}
	}
	return nil
}

func (p *MemoryPersistence) getExerciseById(id int64) *Exercise {
	for i := 0; i < len(p.exercises); i++ {
		if p.exercises[i].ID == id {
			return &p.exercises[i]
		}
	}

	return nil
}

func (p *MemoryPersistence) getAllExercises() []Exercise {
	return p.exercises
}
