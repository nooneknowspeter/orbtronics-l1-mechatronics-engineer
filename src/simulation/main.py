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
        self.temperature = random.randint(18, 35)
        self.humidity = random.randint(40, 80)


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