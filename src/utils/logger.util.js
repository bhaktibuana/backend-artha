const moment = require("moment");
const path = require("path");
const fs = require("fs");

const logFormat = (tokens, req, res) => {
  const date = `[${tokens.date(req, res, "web")}]`;
  const url = `"${tokens.method(req, res)} ${tokens.url(
    req,
    res
  )} HTTP/${tokens["http-version"](req, res)}"`;
  const responseTime = `${tokens["response-time"](req, res)}ms`;
  const ipAddress = `${tokens["remote-addr"](req, res)}`;
  const userAgent = `"${tokens["user-agent"](req, res)}"`;
  return [
    date,
    url,
    tokens.status(req, res),
    tokens.res(req, res, "content-length"),
    "-",
    responseTime,
    ipAddress,
    userAgent,
  ].join(" ");
};

const logOptions = () => {
  const today = moment();
  const date = today.format("YYYYMMDD");
  const sevenDaysAgo = today.subtract(7, "days").format("YYYYMMDD");
  const appDir = process.cwd();
  const logDir = path.join(appDir, "./logs");

  if (!fs.existsSync(logDir)) {
    fs.mkdirSync(logDir);
  } else {
    fs.readdir(logDir, (error, files) => {
      if (error) return;
      files.forEach((file) => {
        const fileDate = file.split("-")[1].split(".")[0];
        if (
          moment(fileDate, "YYYYMMDD").isSameOrBefore(
            moment(sevenDaysAgo, "YYYYMMDD")
          )
        )
          fs.unlink(path.join(logDir, file), (error) => {
            return;
          });
      });
    });
  }

  return {
    stream: fs.createWriteStream(path.join(logDir, `logfile-${date}.log`), {
      flags: "a",
    }),
  };
};

module.exports = {
  logFormat,
  logOptions,
};
