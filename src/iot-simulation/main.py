import random
from time import sleep


class SimulationData:
    temperature: int = 0
    humidity: int = 0

    def __init__(self) -> None:
        None

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