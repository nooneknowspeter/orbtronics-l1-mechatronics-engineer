# TODO

- [ ] transfer main README.md hardware information to cad/README.md

## Part 2 - Software

- [ ] Python script
  - [x] Simulates sensor readings for temperature (18–35°C) and humidity (40–80%)
  - [x] Sends readings every 5–10 seconds to a backend API
  - [ ] Occasionally sends “abnormal” readings to test alerts

- [x] Backend API (FastAPI or Flask preferred)
  - [x] Receives readings and stores them in memory (dict or SQLite)
  - [x] Has endpoint for dashboard to fetch latest readings
  - [ ] Has endpoint to set alert thresholds

- [ ] Dashboard (Simple HTML/JS or React)
  - [ ] Displays latest readings
  - [ ] Shows a red alert if readings exceed thresholds
  - [ ] Allows updating thresholds (sends to backend)

- [ ] Short video demo (max 3 min) showing:
  - [ ] Simulation running
  - [ ] Dashboard updating in real time
  - [ ] Alert triggering

- [ ] README.md – How to run the simulation, backend, and dashboard

- [ ] detail information on software in src/README.md
