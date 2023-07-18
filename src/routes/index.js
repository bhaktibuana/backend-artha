const { Router } = require("express");
const { response, logFormat, logOptions } = require("../utils");
const { xAccessTokenCheck } = require("../middlewares");
const appRoute = require("./app/app.route");
const morgan = require("morgan");

const router = Router();

router.use(morgan(logFormat, logOptions()));
router.use("/api", xAccessTokenCheck, appRoute);

router.use("/:anyRoute", (req, res) => {
  const url = `${req.protocol}://${req.headers.host}${req.originalUrl}`;
  response(`URL not found for: ${url}`, 404, res);
});

router.use("/", (req, res) => {
  const url = `${req.protocol}://${req.headers.host}`;
  response("Artha API", 200, res, { url });
});

module.exports = router;
