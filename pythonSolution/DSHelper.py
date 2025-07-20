# Sieves
def sieve(limit: int) -> list[int]:
    isPrime = [True] * (limit + 1)
    isPrime[0], isPrime[1] = False, False  # 0, 1 is not prime

    for i in range(2, int(limit**0.5) + 1):  # All number in range [2:sqrt(limit)]
        if isPrime[i]:
            for j in range(i * i, limit + 1, i):  # All (i^2) + (i*N) is not a prime
                isPrime[j] = False

    return [i for i, prime in enumerate(isPrime) if prime]


# Linked List
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# D-Link
class DoublyLinkedNode:
    def __init__(self, key, val):
        self.key = key
        self.val = val
        self.prev = None
        self.next = None


# Build a linked list: 1 -> 2 -> 3 -> None
head = ListNode(1, ListNode(2, ListNode(3)))

# Traverse and print the linked li:
node = head
while node:
    print(node.val, end=" -> ")
    node = node.next
print("None")  # Output: 1 -> 2 -> 3 -> None


# Trie
class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False  # Marks end of a complete word


class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for char in word:
            if char not in node.children:
                node.children[char] = TrieNode()
            node = node.children[char]
        node.is_end = True

    def search(self, word):
        node = self.root
        for char in word:
            if char not in node.children:
                return False
            node = node.children[char]
        return node.is_end


# Fenwick Tree
class FenwickTree:
    def __init__(self, size):
        self.tree = [0] * (size + 1)

    def update(self, index, delta):
        index += 1
        while index < len(self.tree):
            self.tree[index] += delta
            index += index & -index

    def query(self, index) -> int:
        index += 1
        res = 0
        while index > 0:
            res += self.tree[index]
            index -= index & -index
        return res


# Disjoint set
class DisjointSet:
    def __init__(self, n):
        self.parent = [i for i in range(n)]  # node is it's own parent
        self.rank = [0] * n  # rank of "parent" - higher bigger

    def find(self, x):
        # compression for better performance
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x, y):
        xroot, yroot = self.find(x), self.find(y)
        if xroot == yroot:
            return

        # Attach smaller tree into bigger tree
        if self.rank[xroot] > self.rank[yroot]:
            self.parent[yroot] = xroot
        elif self.rank[xroot] < self.rank[yroot]:
            self.parent[xroot] = yroot
        else:
            self.parent[yroot] = xroot
            self.rank[xroot] += 1


# Stack
# Python's list can serve as a stack:
stack = []
stack.append(1)  # push
stack.append(3)
print("Stack pop:", stack.pop())  # pop (Output: 3)


# Queue
from collections import deque

# Using collections.deque for efficient FIFO operations:
queue = deque()
queue.append(1)  # enqueue
print("Queue popleft:", queue.popleft())  # dequeue (Output: 1)


# Graph as Adj list
# Graph as an adjacency list (dictionary):
from collections import defaultdict

graph = defaultdict(list)
graph = {1: [2, 3], 2: [1, 4], 3: [1], 4: [2]}
print("Graph:", graph)


# KV Iteration
dictionary = dict()
dictionary["some"] = "thing"
for k, v in dictionary.items():
    pass

# min/max function
someDict = defaultdict(list)
min(someDict)  # Return the lowest key
min(someDict, key=someDict.get)  # Lowest value

# OrderedDict
from collections import OrderedDict

od = OrderedDict()
od["chris"] = 2
od["kek"] = 1
# Iterate "chris" > "kek"
last_key = next(reversed(od))  # Get last in ordDict


# check for digits
someStr = "sada1322"
if any(char.isDigit() for char in someStr):
    pass


# Min Heap
import heapq

# Create a min-heap and push/pop elements.
heap = [5, 7, 9, 1, 3]
heapq.heapify(heap)  # Transform list into a heap.
print("Heap pop:", heapq.heappop(heap))  # Output: 1 (smallest element)

pq = []
heapq.heappush(pq, 1)  # This works too


# Unicode
ord("A")
chr(123)  # reverse

"""
String Ops:
- Change case (e.g., upper(), lower())
- Trim whitespace (strip(), lstrip(), rstrip())
- Split and join sequences of strings (split(), join())
- Search and replace substrings (find(), replace())
- Check string prefixes/suffixes (startswith(), endswith())
- Format strings using .format() or f-strings
- isnumeric() to check number
- can reversed() too
"""


# Sorting
wtf = [[1, 1], [2, 3]]
wtf.sort(key=lambda x: x[0])  # Lambda x could be anything, even an instance
wtf.sort(key=lambda x: (x[0], x[1]))
wtf = sorted(wtf, key=lambda x: x[0])  # Make new

# Sorted List
from sortedcontainers import SortedList

wtf = SortedList([1, 3, 2, 3, 4])
wtf.add(123)  # this gets sorted automatically
# wtf.insert()


"""
Searching
"""
# bisect_left(a, x)	Find leftmost insertion index for x in a
# bisect_right(a, x) Find rightmost insertion index for x in a
# insort_left(a, x)	Insert x at leftmost position in a
# insort_right(a, x) Insert x at rightmost position in a
import bisect

arr = [1, 2, 3, 3, 4, 5, 5, 5, 6, 7]
numbers = [1, 2, 4, 4, 5]
x = 4

# Position x1 < x <= x2 - leftmost of equals elem
left_index = bisect.bisect_left(numbers, x, lo=3, hi=4)
# Position x1 <= x < x2 - rightmost of equals elem
right_index = bisect.bisect_right(numbers, x, lo=3, hi=4)


# Count
arr = [1, 2, 3]
arr.count(2)

# It is different from default dict though
# cword IS A special dict with missing keys return 0
from collections import Counter

word = "asdasdcvs"
cword = Counter(word)

# Language LRU cache
#  @lru_cache(None)  # memoization to avoid recomputation
# nonlocal res  # Declare that 'res' is from the outer scope


# Starter func
if __name__ == "__main__":
    cases = []

    for c in cases:
        print()
