package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// rest endpoints:
// 1. create exercise
// 2. get all exercises
// 3. get exercise by id

// 4. perform exercise
// 5. get all completed by exercise id

func apiGetExercise(c *gin.Context, p PersistanceProvider) {
	// gin does not support regex for url param, so we are not able to restrict id to integers
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}

	exercise := p.getExerciseById(id)

	if exercise == nil {
		c.JSON(404, gin.H{"code": "ELEMENT_NOT_FOUND", "message": "Exercise not found"})
	} else {
		c.IndentedJSON(http.StatusOK, exercise)
	}
}

func apiGetAllExercises(c *gin.Context, p PersistanceProvider) {
	data := p.getAllExercises()

	if data == nil {
		c.IndentedJSON(http.StatusOK, Response{Status: "empty data"})
	} else {
		c.IndentedJSON(http.StatusOK, data)
	}
}

func apiAddExercise(c *gin.Context, p PersistanceProvider) {
	var newExercise ExerciseCreate

	if err := c.BindJSON(&newExercise); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p.addExercise(newExercise)

	c.IndentedJSON(http.StatusOK, Response{Status: "ok"})
}

func apiDeleteExercise(c *gin.Context, p PersistanceProvider) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}

	exercises := p.getAllExercises()

	// Find the index of the element to delete
	indexToDelete := -1
	for i, exercise := range exercises {
		if exercise.ID == id {
			indexToDelete = i + 1
			break
		}
	}

	if indexToDelete == -1 {
		c.JSON(404, gin.H{"code": "ELEMENT_NOT_FOUND", "message": "Exercise not found"})
	} else {
		p.deleteExercise(int64(indexToDelete))
		c.IndentedJSON(http.StatusOK, Response{Status: "ok"})
	}
}

func apiUpdateExercise(c *gin.Context, p PersistanceProvider) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	var updatedExercise ExerciseUpdate

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}

	if err := c.BindJSON(&updatedExercise); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := p.updateExerciseById(id, updatedExercise)

	if result != nil {
		c.IndentedJSON(http.StatusOK, result)
	} else {
		c.JSON(404, gin.H{"code": "ELEMENT_NOT_FOUND", "message": "Exercise not found"})
	}
}
