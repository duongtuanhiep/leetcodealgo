"""
Linked List
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


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

"""
Stack
"""
# Python's list can serve as a stack:
stack = []
stack.append(1)  # push
stack.append(3)
print("Stack pop:", stack.pop())  # pop (Output: 3)


"""
Queue
"""
from collections import deque

# Using collections.deque for efficient FIFO operations:
queue = deque()
queue.append(1)  # enqueue
print("Queue popleft:", queue.popleft())  # dequeue (Output: 1)


"""
Graph as a dict()
"""
# Graph as an adjacency list (dictionary):
from collections import defaultdict

graph = defaultdict(list)
graph = {1: [2, 3], 2: [1, 4], 3: [1], 4: [2]}
print("Graph:", graph)


"""
Min Heap
"""
import heapq

# Create a min-heap and push/pop elements.
heap = [5, 7, 9, 1, 3]
heapq.heapify(heap)  # Transform list into a heap.
print("Heap pop:", heapq.heappop(heap))  # Output: 1 (smallest element)

pq = []
heapq.heappush(pq, 1)  # This works too


"""
Unicode 
"""
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
- And more utilities via the string module
"""

"""
Sorting
"""
wtf = [[1, 1], [2, 3]]
wtf.sort(key=lambda x: x[0])  # Lambda x could be anything, even an instance
wtf.sort(key=lambda x: (x[0], x[1]))
wtf = sorted(wtf, key=lambda x: x[0])  # Make new


"""
Searching
"""
# bisect_left(a, x)	Find leftmost insertion index for x in a
# bisect_right(a, x)	Find rightmost insertion index for x in a
# insort_left(a, x)	Insert x at leftmost position in a
# insort_right(a, x)	Insert x at rightmost position in a
import bisect

arr = [1, 2, 3, 3, 4, 5, 5, 5, 6, 7]
numbers = [1, 2, 4, 4, 5]
x = 4

# Position x1 < x <= x2 - leftmost of equals elem
left_index = bisect.bisect_left(numbers, x)
# Position x1 <= x < x2 - rightmost of equals elem
right_index = bisect.bisect_right(numbers, x)


"""
Count
"""
arr = [1, 2, 3]
arr.count(2)

"""
KV Iteration
"""
dictionary = dict()
dictionary["some"] = "thing"

for k, v in dictionary.items():
    pass

# nonlocal res  # Declare that 'res' is from the outer scope


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


# min/max function
someDict = defaultdict(list)
min(someDict)  # Return the lowest key
min(someDict, key=someDict.get)  # Lowest value

# check for digits
someStr = "sada1322"
if any(char.isDigit() for char in someStr):
    pass

# Starter func
if __name__ == "__main__":
    cases = []

    for c in cases:
        print()

# Language LRU cache
#  @lru_cache(None)  # memoization to avoid recomputation
