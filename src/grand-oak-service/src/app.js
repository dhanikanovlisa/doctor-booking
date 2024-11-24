const express = require("express");
const app = express();
const port = 3001;

app.use(express.json());

app.get("/", (_, res) => {
    res.json({message: 'Welcome To Grand Oak Hospital Service'})
})

// Available appointment slots
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
});

app.listen(port, () => {
    console.log(`Grand Oak Service running on http://localhost:${port}`);
});