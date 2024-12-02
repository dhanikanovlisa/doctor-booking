
const express = require("express");
const Joi = require("joi");
const slots = require("../data/slots");

const router = express.Router();

// GET available-slots from data "booked" : false
router.get("/available-slots", (_, res) => {
  const availableSlots = slots.filter((slot) => !slot.booked);
  res.json(availableSlots);
});

// POST booking
router.post("/book/:id", (req, res) => {

  const slotId = parseInt(req.params.id, 10);
  const slot = slots.find((slot) => slot.id === slotId);

  if (!slot) {
    return res.status(404).json({ message: "Slot not found" });
  }
  if (slot.booked) {
    return res.status(400).json({ message: "Slot is already booked" });
  }

  slot.booked = true;

  res.json({ message: `Appointment booked successfully`, slot });
});

module.exports = router;
