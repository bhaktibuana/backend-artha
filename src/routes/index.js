const { Router } = require("express");

const router = Router();

router.use("/", (req, res) => {
  res.status(200).json({
    message: "Artha API",
  });
});

module.exports = router;
