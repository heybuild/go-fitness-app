package main

type PersistanceProvider interface {
	addExercise(exercide ExerciseCreate) bool
	deleteExercise(id int64) bool
	updateExerciseById(id int64, exercise ExerciseUpdate) *Exercise
	getExerciseById(id int64) *Exercise
	getAllExercises() []Exercise
}
