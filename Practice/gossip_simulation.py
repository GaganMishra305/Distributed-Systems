import random

class GossipNode:
    def __init__(self, node_id: int):
        self.node_id = node_id
        self.state = {}

    def update_state(self, key, value):
        self.state[key] = value

    def gossip_round(self, peers):
        if not peers:
            return
        peer = random.choice(peers)
        peer.state.update(self.state)
