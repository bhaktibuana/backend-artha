const { config: dotenvConfig } = require("dotenv");
dotenvConfig();

const serverPort = process.env.SERVER_PORT;
const jwtSecretKey = process.env.JWT_SECRET_KEY;
const jwtExpiredTime = process.env.JWT_EXPIRED_TIME;
const dbHost = process.env.DB_HOST;
const dbUsername = process.env.DB_USER;
const dbPassword = process.env.DB_PASSWORD;
const dbName = process.env.DB_NAME;
const dbPort = process.env.DB_PORT;
const smtpPort = process.env.SMTP_PORT;
const smtpHost = process.env.SMTP_HOST;
const smtpUsername = process.env.SMTP_USERNAME;
const smtpPassword = process.env.SMTP_PASSWORD;

const config =
  process.env.NODE_ENV === "development"
    ? {
        nodeEnv: "development",
        serverPort: 3001,
        jwtSecretKey: jwtSecretKey !== undefined ? jwtSecretKey : "",
        jwtExpiredTime: jwtExpiredTime !== undefined ? jwtExpiredTime : "",
        xAccessToken: "api-artha",
        xAccessTokenTest: "@p!.@rth4",
        dbHost: dbHost !== undefined ? dbHost : "",
        dbUsername: dbUsername !== undefined ? dbUsername : "",
        dbPassword: dbPassword !== undefined ? dbPassword : "",
        dbName: dbName !== undefined ? dbName : "",
        dbPort: 3306,
        smtpPort: smtpPort !== undefined ? smtpPort : "",
        smtpHost: smtpHost !== undefined ? smtpHost : "",
        smtpUsername: smtpUsername !== undefined ? smtpUsername : "",
        smtpPassword: smtpPassword !== undefined ? smtpPassword : "",
      }
    : {
        nodeEnv: "production",
        serverPort: serverPort !== undefined ? parseInt(serverPort) : "",
        jwtSecretKey: jwtSecretKey !== undefined ? jwtSecretKey : "",
        jwtExpiredTime: jwtExpiredTime !== undefined ? jwtExpiredTime : "",
        xAccessToken: "api-artha",
        xAccessTokenTest: "@p!.@rth4",
        dbHost: dbHost !== undefined ? dbHost : "",
        dbUsername: dbUsername !== undefined ? dbUsername : "",
        dbPassword: dbPassword !== undefined ? dbPassword : "",
        dbName: dbName !== undefined ? dbName : "",
        dbPort: dbPort !== undefined ? parseInt(dbPort) : "",
        smtpPort: smtpPort !== undefined ? smtpPort : "",
        smtpHost: smtpHost !== undefined ? smtpHost : "",
        smtpUsername: smtpUsername !== undefined ? smtpUsername : "",
        smtpPassword: smtpPassword !== undefined ? smtpPassword : "",
      };

module.exports = { config };
