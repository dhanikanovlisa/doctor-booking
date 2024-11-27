package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type Slot struct {
	ID        int    `json:"id"`
	Time      string `json:"time"`
	Doctor    string `json:"doctor"`
	Specialty string `json:"specialty"`
}

// Mock data for available slots
var slots = []Slot{
	{ID: 1, Time: "09:00", Doctor: "Dr. Watson", Specialty: "Cardiology"},
	{ID: 2, Time: "10:00", Doctor: "Dr. Watson", Specialty: "Cardiology"},
	{ID: 3, Time: "11:00", Doctor: "Dr. Watson", Specialty: "Cardiology"},
	{ID: 4, Time: "12:00", Doctor: "Dr. Brown", Specialty: "Dermatology"},
	{ID: 5, Time: "11:00", Doctor: "Dr. White", Specialty: "Neurology"},
	{ID: 6, Time: "09:00", Doctor: "Dr. Green", Specialty: "Neurology"},
	{ID: 7, Time: "11:00", Doctor: "Dr. Black", Specialty: "Dermatology"},
	{ID: 8, Time: "10:00", Doctor: "Dr. Blue", Specialty: "Gynecology"},
}

func main() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default())

	// Welcome message handler
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome To Pine Valley Hospital Service"})
	})

	// Get available slots
	router.GET("/available-slots", func(c *gin.Context) {
		c.JSON(http.StatusOK, slots)
	})

	// Book a doctor slot
	router.POST("/book/:id", func(c *gin.Context) {
		slotID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid slot ID"})
			return
		}

		for i, slot := range slots {
			if slot.ID == slotID {
				slots = append(slots[:i], slots[i+1:]...)

				c.JSON(http.StatusOK, gin.H{
					"message": "Appointment booked successfully",
					"data":    slot,
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"message": "Slot ID not found"})
	})

	// Start server
	port := ":3002"
	log.Printf("Pine Valley Service running on http://localhost%s\n", port)
	router.Run(port)
}
