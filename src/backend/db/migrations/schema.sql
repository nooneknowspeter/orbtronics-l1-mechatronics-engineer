CREATE TABLE IF NOT EXISTS devices (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    temperature REAL,
    humidity REAL
);