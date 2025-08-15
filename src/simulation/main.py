import random
from time import sleep


class SimulationData:
    device_name: str = "device001"
    temperature: float = 0
    humidity: float = 0

    def __init__(self) -> None:
        None

    def randomValue(self, start: int, end: int, roundNDigits: int = 2) -> float:
        return round(random.uniform(start, end), roundNDigits)

    def updateValues(self) -> None:
        self.humidity = self.randomValue(18, 35, 2)
        self.temperature = self.randomValue(40, 80, 2)


def simulation() -> None:
    data = SimulationData()

    while True:
        data.updateValues()
        print(
            "Temperature:",
            f"{data.temperature}°C",
            "|",
            "Humidity:",
            f"{data.humidity}%",
        )
        sleep(5)


if __name__ == "__main__":
    simulation()