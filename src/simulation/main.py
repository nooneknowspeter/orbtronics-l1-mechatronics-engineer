import os
import random
from time import sleep

import dotenv
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


def simulation(data, server_url) -> None:
    while True:
        data.updateValues()
        data.postDataToServer(server_url=server_url)

        print(data)

        sleep(5)


def loadEnvironmentFile():
    parent_env_file = "../../.env"
    cwd_env_file = ".env"

    try:
        if os.path.exists(parent_env_file):
            print("loading parent .env file")
            dotenv.load_dotenv(dotenv_path=parent_env_file)
        elif os.path.exists(cwd_env_file):
            print("loading cwd .env file")
            dotenv.load_dotenv()
        else:
            raise FileNotFoundError(".env files not found")
    except Exception as e:
        print(e)


if __name__ == "__main__":
    dotenv.load_dotenv()

    server_url = os.getenv("SERVER_URL")
    if server_url is None:
        print("environment variable SERVER_URL not defined")
    else:
        data = SimulationData()

        simulation(data, server_url)
