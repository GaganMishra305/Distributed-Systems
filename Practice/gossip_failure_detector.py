import time

class HeartbeatDetector:
    def __init__(self, timeout: float = 3.0):
        self.timeout = timeout
        self.last_heartbeats = {}

    def heartbeat(self, node_id: str):
        self.last_heartbeats[node_id] = time.time()

    def check_failures(self):
        now = time.time()
        return [node for node, ts in self.last_heartbeats.items() if now - ts > self.timeout]
