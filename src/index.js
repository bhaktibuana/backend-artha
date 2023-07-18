const app = require("./app");
const { config } = require("./configs");

app.listen(config.serverPort, () => {
  console.log("App is running on port", config.serverPort);
});
