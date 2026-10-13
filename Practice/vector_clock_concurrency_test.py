import unittest
from vector_clock_example import VectorClock

class TestVectorClockConcurrency(unittest.TestCase):
    def test_causality(self):
        c1 = VectorClock(0, 2)
        c2 = VectorClock(1, 2)
        e1 = c1.send_event()
        c2.receive_event(e1)
        self.assertTrue(c1.clock[0] < c2.clock[0])

if __name__ == "__main__":
    unittest.main()
