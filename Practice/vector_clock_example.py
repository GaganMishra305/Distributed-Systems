class VectorClock:
    def __init__(self, process_id: int, total_nodes: int):
        self.pid = process_id
        self.clock = [0] * total_nodes

    def increment(self):
        self.clock[self.pid] += 1

    def send_event(self):
        self.increment()
        return list(self.clock)

    def receive_event(self, msg_clock):
        for i in range(len(self.clock)):
            self.clock[i] = max(self.clock[i], msg_clock[i])
        self.clock[self.pid] += 1
