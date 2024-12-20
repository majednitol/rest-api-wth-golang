package controllers

import (
	"net/http"

	"github.com/counter/middleware"
	models "github.com/counter/model"
	"github.com/counter/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CounterController struct {
	Model *models.CounterModel
}

// GetCounters fetches all counters from the database
func (cc *CounterController) GetCounters(c *gin.Context) {
	counters, err := cc.Model.FetchAllCounters()
	if err != nil {
		c.Error(&middleware.AppError{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch counters",
		})
		return
	}
	c.JSON(http.StatusOK, counters)
}

// CreateCounter creates a new counter in the database
func (cc *CounterController) CreateCounter(c *gin.Context) {
	counter := models.Counter{Value: 0}
	id, err := cc.Model.CreateCounter(counter)
	if err != nil {
		c.Error(&middleware.AppError{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create counter",
		})
		return
	}
	counter.ID = id

	// Broadcast updated counters to WebSocket clients
	counters, fetchErr := cc.Model.FetchAllCounters()
	if fetchErr == nil {
		services.BroadcastCounters(counters)
	}

	c.JSON(http.StatusOK, counter)
}

// UpdateCounter updates the value of an existing counter
func (cc *CounterController) UpdateCounter(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(&middleware.AppError{
			Code:    http.StatusBadRequest,
			Message: "Invalid counter ID",
		})
		return
	}

	var body struct {
		Value int `json:"value"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(&middleware.AppError{
			Code:    http.StatusBadRequest,
			Message: "Invalid JSON",
		})
		return
	}

	if err := cc.Model.UpdateCounter(id, body.Value); err != nil {
		c.Error(&middleware.AppError{
			Code:    http.StatusNotFound,
			Message: "Counter not found",
		})
		return
	}

	// Broadcast updated counters to WebSocket clients
	counters, fetchErr := cc.Model.FetchAllCounters()
	if fetchErr == nil {
		services.BroadcastCounters(counters)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Counter updated"})
}

// DeleteCounter deletes a counter by its ID
func (cc *CounterController) DeleteCounter(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(&middleware.AppError{
			Code:    http.StatusBadRequest,
			Message: "Invalid counter ID",
		})
		return
	}

	if err := cc.Model.DeleteCounter(id); err != nil {
		c.Error(&middleware.AppError{
			Code:    http.StatusNotFound,
			Message: "Counter not found",
		})
		return
	}

	// Broadcast updated counters to WebSocket clients
	counters, fetchErr := cc.Model.FetchAllCounters()
	if fetchErr == nil {
		services.BroadcastCounters(counters)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Counter deleted"})
}
