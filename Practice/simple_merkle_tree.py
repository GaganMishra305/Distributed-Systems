import hashlib

class MerkleTree:
    def __init__(self, leaves):
        self.leaves = [hashlib.sha256(l.encode()).hexdigest() for l in leaves]
        self.root = self._build_tree(self.leaves)

    def _build_tree(self, nodes):
        if len(nodes) == 1:
            return nodes[0]
        new_level = []
        for i in range(0, len(nodes), 2):
            left = nodes[i]
            right = nodes[i + 1] if i + 1 < len(nodes) else left
            combined = hashlib.sha256((left + right).encode()).hexdigest()
            new_level.append(combined)
        return self._build_tree(new_level)
