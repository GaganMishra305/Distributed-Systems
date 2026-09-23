class Coordinator:
    def __init__(self, participants):
        self.participants = participants

    def execute_transaction(self, tx_id):
        # Phase 1: Prepare
        votes = [p.prepare(tx_id) for p in self.participants]
        if all(votes):
            # Phase 2: Commit
            for p in self.participants:
                p.commit(tx_id)
            return "COMMITTED"
        else:
            # Phase 2: Abort
            for p in self.participants:
                p.abort(tx_id)
            return "ABORTED"
