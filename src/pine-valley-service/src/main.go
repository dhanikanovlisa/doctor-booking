package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Slot struct {
	ID     int    `json:"id"`
	Time   string `json:"time"`
	Doctor string `json:"doctor"`
}

type BookingRequest struct {
	SlotID      int    `json:"slotId"`
	PatientName string `json:"patientName"`
}

func main() {
	// Available appointment slots
	slots := []Slot{
		{ID: 1, Time: "9:00 AM", Doctor: "Dr. Watson"},
		{ID: 2, Time: "10:30 AM", Doctor: "Dr. Johnson"},
	}

	// Welcome message handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"message": "Welcome To Pine Valley Hospital Service"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Available slots handler
	http.HandleFunc("/available-slots", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(slots)
	})

	// Book appointment handler
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var booking BookingRequest
		if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if booking.SlotID == 0 || booking.PatientName == "" {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		response := map[string]string{
			"message": "Appointment booked successfully with ID " +
				string(rune(booking.SlotID)) + " for " + booking.PatientName,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Start the server
	port := ":3002"
	log.Printf("Pine Valley Service running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
