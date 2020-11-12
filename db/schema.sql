-- Create tables.
DROP TABLE IF EXISTS plants;
CREATE TABLE plants
(
 id SERIAL PRIMARY KEY,
 soilMoistureLevel REAL CHECK (soilMoistureLevel >= 0 AND soilMoistureLevel <= 1),
 soilDataTimestamp TIMESTAMP default current_timestamp
);

-- Insert demo data.
INSERT INTO plants (soilMoistureLevel) VALUES (0.1);
INSERT INTO plants (soilMoistureLevel) VALUES (0.15);
INSERT INTO plants (soilMoistureLevel) VALUES (0.2);
INSERT INTO plants (soilMoistureLevel) VALUES (0.25);
INSERT INTO plants (soilMoistureLevel) VALUES (0.3);
