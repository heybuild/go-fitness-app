package main

import (
	"github.com/gin-gonic/gin"
)

// entities
// - exercises (id, name, description)
// - completed (id, datum, reps, kg, exercise_id)

// rest endpoints:
// 1. create exercise
// 2. get all exercises
// 3. get exercise by id
// 4. perform exercise
// 5. get all completed by exercise id

func main() {
	provider := &MemoryPersistence{}

	hello_var := Hello("testing")
	hello_var.testing()

	router := gin.Default()
	router.GET("/exercises", func(c *gin.Context) {
		apiGetAllExercises(c, provider)
	})
	router.POST("/exercises", func(c *gin.Context) {
		apiAddExercise(c, provider)
	})
	router.GET("/exercises/:id", func(c *gin.Context) {
		apiGetExercise(c, provider)
	})
	router.DELETE("/exercises/:id", func(c *gin.Context) {
		apiDeleteExercise(c, provider)
	})
	router.PATCH("/exercises/:id", func(c *gin.Context) {
		apiUpdateExercise(c, provider)
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	router.Run("localhost:8080")
}
