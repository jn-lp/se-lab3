const { Client } = require("./plants/client");

const client = Client("http://localhost:8080");

(async () => {
  // Scenario 1: Create new plant
  console.log("=== Scenario 1 ===");
  try {
    const newPlant = await client.createPlant();
    console.log("Create plant response:", newPlant);
  } catch (err) {
    console.log(`Problem creating new plant: `, err);
  }

  // Scenario 2: Display plants with low soil moisture level
  let plants;
  console.log("=== Scenario 2 ===");
  try {
    plants = await client.listPlants();
    console.log("Plants with low soil moisture level:");
    console.table(plants);
  } catch (err) {
    console.log(`Problem listing plants: `, err);
  }

  // Scenario 3: Update plants soil moisture level
  console.log("=== Scenario 3 ===");
  const plant = await plants[~~(Math.random() * plants.length)];
  const newLevel = plant.soilMoistureLevel - (Math.random() * 2 - 1);
  try {
    const updatedPlant = await client.updatePlant(
      plant.id,
      newLevel > 0 ? newLevel : -newLevel
    );
    console.log("Update plant response:", updatedPlant);
  } catch (err) {
    console.log(`Problem updating plant: `, err);
  }
})();
