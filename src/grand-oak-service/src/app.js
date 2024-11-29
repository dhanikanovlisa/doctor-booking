
require("dotenv").config();
const express = require("express");
const logger = require("./middleware/logger");
const rootRoutes = require("./routes/index");
const bookingRoutes = require("./routes/booking");


const app = express();
const port = process.env.port || 3001;


//Midlleware
app.use(express.json());
app.use(logger)

//Routes home
app.use('/',rootRoutes);
/*app.get("/", (_, res) => {
    res.json({message: 'Welcome To Grand Oak Hospital Service'})
})*/

//Routes API
app.use("/api", bookingRoutes);
/*
app.get("/available-slots", (req, res) => {
    res.json([
        { id: 1, time: "9:00 AM", doctor: "Dr. Watson" },
        { id: 2, time: "10:30 AM", doctor: "Dr. Johnson" },
    ]);
});

// Book an appointment
app.post("/book", (req, res) => {
    const { slotId, patientName } = req.body;
    if (!slotId || !patientName) {
        return res.status(400).json({ message: "Invalid input" });
    }
    res.json({ message: `Appointment booked successfully with ID ${slotId} for ${patientName}` });
});*/

//Error Handling
app.use((err, req, res, next) => {
    console.error(err.stack);
    res.status(500).json({message : " Server Error"})
})

app.listen(port, () => {
    console.log(`Grand Oak Service running on http://localhost:${port}`);
});