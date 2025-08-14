import random
class SimulationData:
    temperature: int = 0
    humidity: int = 0

    def __init__(self) -> None:
        None

    def updateValues(self) -> None:
        self.temperature = random.randint(18, 35)
        self.humidity = random.randint(40, 80)

