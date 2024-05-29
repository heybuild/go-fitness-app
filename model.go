package main

type Exercise struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ExerciseCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ExerciseUpdate struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type PerformedExercise struct {
	ID       int64    `json:"id"`
	Date     string   `json:"date"`
	Reps     int      `json:"reps"`
	Weight   float32  `json:"weight"`
	Exercise Exercise `json:"exercise"`
}

type Response struct {
	Status string `json:"status"`
}
