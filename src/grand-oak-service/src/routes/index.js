
const express = require("express");
const router = express.Router();

router.get("/", (_, res) => {
  res.json({ message: "Welcome to Grand Oak Hospital Service" });
});

module.exports = router;
