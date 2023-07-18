const jwt = require("jsonwebtoken");
const { config } = require("../../configs");
const { response } = require("../../utils");

const xAccessTokenCheck = (req, res, next) => {
  if (req.headers["x-access-token"] === undefined) {
    response("Forbidden", 403, res);
  } else {
    const splitToken = req.headers["x-access-token"].split(" ");

    if (splitToken.length !== 2 || splitToken[0] !== "Bearer") {
      response("Wrong x-access-token format", 400, res);
    } else if (splitToken[1] === config.xAccessTokenTest) {
      next();
    } else {
      jwt.verify(
        splitToken[1],
        config.jwtSecretKey,
        { algorithms: ["HS256"] },
        (error, payload) => {
          if (error && error.name === "TokenExpiredError") {
            response("Forbidden: Token expired", 403, res);
          } else if (error) {
            response("Forbidden: Invalid token", 403, res);
          } else {
            const tokenPayload =
              typeof payload === "object"
                ? payload["x-access-token"]
                : undefined;

            tokenPayload === config.xAccessToken
              ? next()
              : response("Invalid x-access-token", 403, res);
          }
        }
      );
    }
  }
};

module.exports = {
  xAccessTokenCheck,
};
