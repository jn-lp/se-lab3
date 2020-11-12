const http = require("../common/http");

const Client = (baseUrl) => (
  (client = http.Client(baseUrl)),
  {
    listPlants: () => client.get("/plants"),
    createPlant: () => client.post("/plants"),
    updatePlant: (id, soilMoistureLevel) =>
      client.patch(`/plants`, { id, soilMoistureLevel }),
  }
);

module.exports = { Client };
