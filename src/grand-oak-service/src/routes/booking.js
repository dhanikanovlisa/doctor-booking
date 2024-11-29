
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
router.post("/booking", (req, res) => {
  const schema = Joi.object({
    slotId: Joi.number().required(),
    patientName: Joi.string().min(3).required(),
  });

  const { error } = schema.validate(req.body);
  if (error) {
    return res.status(400).json({ message: error.details[0].message });
  }

  const { slotId, patientName } = req.body;
  const slot = slots.find((slot) => slot.id === slotId);

  if (!slot) {
    return res.status(404).json({ message: "Slot not found" });
  }
  if (slot.booked) {
    return res.status(400).json({ message: "Slot is already booked" });
  }

  slot.booked = true;
  slot.patientName = patientName;

  res.json({ message: `Appointment booked successfully`, slot });
});

module.exports = router;
