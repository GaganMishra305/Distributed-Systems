class RaftNode:
    def __init__(self, node_id, peers):
        self.node_id = node_id
        self.peers = peers
        self.current_term = 0
        self.voted_for = None
        self.state = "FOLLOWER"

    def request_vote(self, term, candidate_id):
        if term > self.current_term:
            self.current_term = term
            self.voted_for = None
            self.state = "FOLLOWER"

        if self.voted_for is None or self.voted_for == candidate_id:
            self.voted_for = candidate_id
            return True
        return False
