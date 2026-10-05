import unittest
from consistent_hashing import ConsistentHashRing

class TestConsistentHash(unittest.TestCase):
    def test_key_distribution(self):
        ring = ConsistentHashRing(replicas=50)
        ring.add_node("node-1")
        ring.add_node("node-2")
        ring.add_node("node-3")
        nodes = set(ring.get_node(f"user_{i}") for i in range(100))
        self.assertEqual(len(nodes), 3)

if __name__ == "__main__":
    unittest.main()
