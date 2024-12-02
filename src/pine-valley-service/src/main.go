package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Slot struct {
	ID        int    `json:"id"`
	Time      string `json:"time"`
	Doctor    string `json:"doctor"`
	Specialty string `json:"specialty"`
	Booked    bool   `json:"booked"`
}

// Mock data for available slots
var slots = []Slot{
	{ID: 1, Time: "09:00 AM", Doctor: "Dr. Alice", Specialty: "Pediatrics", Booked: false},
	{ID: 2, Time: "10:30 AM", Doctor: "Dr. Bob", Specialty: "Orthopedics", Booked: false},
	{ID: 3, Time: "11:00 AM", Doctor: "Dr. Charlie", Specialty: "Cardiology", Booked: false},
	{ID: 4, Time: "01:00 PM", Doctor: "Dr. Diana", Specialty: "Neurology", Booked: false},
	{ID: 5, Time: "02:00 PM", Doctor: "Dr. Evelyn", Specialty: "Dermatology", Booked: false},
	{ID: 6, Time: "03:00 PM", Doctor: "Dr. Frank", Specialty: "Gynecology", Booked: false},
	{ID: 7, Time: "04:00 PM", Doctor: "Dr. Grace", Specialty: "Psychiatry", Booked: false},
	{ID: 8, Time: "05:30 PM", Doctor: "Dr. Henry", Specialty: "ENT", Booked: false},
}

func main() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default())

	// Group routes under /pine-valley
	pineValley := router.Group("/pine-valley")
	{
		// Welcome message handler
		pineValley.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome To Pine Valley Hospital Service"})
		})

		// Get available slots
		pineValley.GET("/available-slots", func(c *gin.Context) {
			var availableSlots []Slot
			for _, slot := range slots {
				if !slot.Booked {
					availableSlots = append(availableSlots, slot)
				}
			}

			// Respond with the list of available slots
			c.JSON(http.StatusOK, availableSlots)
		})

		// Book a doctor slot
		pineValley.POST("/book/:id", func(c *gin.Context) {
			slotID, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid slot ID"})
				return
			}

			for i, slot := range slots {
				if slot.ID == slotID {
					if slot.Booked {
						c.JSON(http.StatusPreconditionFailed, gin.H{
							"message": "Appointment is already booked",
						})
						return
					}

					// Booked
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
	}

	// Start server
	port := ":3002"
	log.Printf("Pine Valley Service running on http://localhost%s/pine-valley\n", port)
	router.Run(port)
}
