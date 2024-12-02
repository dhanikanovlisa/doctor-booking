require("dotenv").config(); // Load environment variables from a .env file
const express = require("express");
const axios = require("axios");
const app = express();

const port = 8080; 
const pineValleyService = process.env.PINE_VALLEY_SERVICE_URL ||'http://pinevalleyservice/pine-valley';
const grandOakService = process.env.GRAND_OAK_SERVICE_URL || 'http://grandoakservice/grand-oak';

app.use(express.json());

// Root Endpoint
app.get("/doctor-booking", (_, res) => {
    res.json({ message: "Welcome to the Doctor Booking Service" });
});

// Fetch available slots from both services
app.get("/doctor-booking/available-slots", async (req, res) => {
    try {``
        // Fetch data from both hospital services concurrently
        const [pineValleyResponse, grandOakResponse] = await Promise.all([
            axios.get(`${pineValleyService}/available-slots`),
            axios.get(`${grandOakService}/available-slots`),
        ]);

        // Combine the data from both services
        const combinedSlots = [
            ...pineValleyResponse.data.map((slot) => ({
                ...slot,
                source: "Pine Valley Hospital",
            })),
            ...grandOakResponse.data.map((slot) => ({
                ...slot,
                source: "Grand Oak Hospital",
            })),
        ];

        res.json(combinedSlots);
    } catch (error) {
        console.error("Error fetching available slots:", error.message);
        res.status(500).json({ message: "Error fetching available slots" });
    }
});

// Book an appointment by forwarding the request to the appropriate service
app.post("/doctor-booking/book", async (req, res) => {
    const { slotId, patientName, hospital } = req.body;

    // Validate input
    if (!slotId || !patientName || !hospital) {
        return res.status(400).json({ message: "Invalid input" });
    }

    // Determine which hospital service to send the request to
    const serviceUrl =
        hospital === "Pine Valley"
            ? `${pineValleyService}/book`
            : hospital === "Grand Oak"
            ? `${grandOakService}/book`
            : null;

    if (!serviceUrl) {
        return res.status(400).json({ message: "Invalid hospital name" });
    }

    try {
        // Forward the booking request to the corresponding service
        const response = await axios.post(serviceUrl, { slotId, patientName });
        res.json(response.data);
    } catch (error) {
        console.error("Error booking appointment:", error.message);
        res.status(500).json({ message: "Error booking appointment" });
    }
});

// Start the service
app.listen(port, () => {
    console.log(`Doctor Booking Service running on http://localhost:${port}`);
    console.log(`URL pine valley service: ${pineValleyService}`)
    console.log(`URL grand oak service: ${grandOakService}`)
});