from collections import OrderedDict
from collections import defaultdict

"""
Currently we can do get in 0(1). 

Key magic is in defaultdict(OrderedDict):
we can have:
- Key as usageCnt
- Valus as orderedDict(K=key,V=1), we need it for the ordering. 
"""


class LFUCache:
    def __init__(self, capacity: int):
        self.keyToVal = dict()
        self.keyToCnt = dict()
        self.cntToKey = defaultdict(OrderedDict)
        self.minUsage = 1
        self.cap = capacity

    def get(self, key: int) -> int:
        if key in self.keyToVal:
            self.incrementCnt(key)
            return self.keyToVal[key]
        return -1

    def put(self, key: int, value: int) -> None:
        if key in self.keyToVal:
            self.incrementCnt(key)
        elif len(self.keyToVal) == self.cap:
            # Remove LRU/LFU keys
            deletedKey, _ = self.cntToKey[self.minUsage].popitem()  # Pop last
            self.keyToVal.pop(deletedKey)
            self.keyToCnt.pop(deletedKey)

        # Adding new key
        if key not in self.keyToVal:
            self.minUsage = 1  # reset min usage
            self.cntToKey[self.minUsage][key] = 1
            self.cntToKey[self.minUsage].move_to_end(key, last=False)
            self.keyToCnt[key] = 1

        # Update key
        self.keyToVal[key] = value

    def incrementCnt(self, key: int) -> None:
        # Retrieve and compute counts
        prevCnt = self.keyToCnt[key]
        newCount = prevCnt + 1

        # remove current key from odr set
        self.cntToKey[prevCnt].move_to_end(key)
        self.cntToKey[prevCnt].popitem(key)

        # increase minCount if required
        if len(self.cntToKey[self.minUsage]) == 0:
            self.minUsage = newCount

        # add current key from new ord set
        self.keyToCnt[key] = newCount
        self.cntToKey[newCount][key] = 1
        self.cntToKey[newCount].move_to_end(key, last=False)


# Your LFUCache object will be instantiated and called as such:
# obj = LFUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
