# TODO

## Part 2 - Software

- [x] simulation
  - [x] Simulates sensor readings for temperature (18–35°C) and humidity (40–80%)
  - [x] Sends readings every 5–10 seconds to a backend API
  - [x] Occasionally sends “abnormal” readings to test alerts
  - [x] dockerfile
  - [x] fix exceptions in simulation, docker

- [x] Backend API (FastAPI or Flask preferred)
  - [x] Receives readings and stores them in memory (dict or SQLite)
  - [x] Has endpoint for dashboard to fetch latest readings
  - [ ] Has endpoint to set alert thresholds
  - [x] dockerfile
  - [ ] setup database

- [x] Dashboard (Simple HTML/JS or React)
  - [ ] ui/ux mockup
  - [x] Displays latest readings
  - [x] Shows a red alert if readings exceed thresholds
  - [ ] Allows updating thresholds (sends to backend)
  - [x] dockerfile

- [x] Short video demo (max 3 min) showing:
  - [x] Simulation running
  - [x] Dashboard updating in real time
  - [x] Alert triggering

- [ ] docs
  - [ ] README.md – How to run the simulation, backend, and dashboard
  - [ ] detail information on software in README.md