import random
from time import sleep

import requests


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

    def postDataToServer(self, server_url) -> int:
        payload = {
            "deviceName": self.device_name,
            "temperature": self.temperature,
            "humidity": self.humidity,
        }

        request = requests.post(url=server_url, json=payload)

        if request.status_code != 201:
            print("failed to send information to server")

    def __str__(self) -> str:
        return str(
            f"""
deviceName: {self.device_name}
temperature: {self.temperature}°C
humidity: {self.humidity}%
"""
        )

    while True:
        data.updateValues()
        print(data)

        sleep(5)


if __name__ == "__main__":
    simulation()