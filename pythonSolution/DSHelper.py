"""
Linked List
"""
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

# Build a linked list: 1 -> 2 -> 3 -> None
head = ListNode(1, ListNode(2, ListNode(3)))

# Traverse and print the linked list:
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
queue.append(1)          # enqueue
print("Queue popleft:", queue.popleft())  # dequeue (Output: 1)

"""
Graph as a dict()
"""
# Graph as an adjacency list (dictionary):
graph = {
    1: [2, 3],
    2: [1, 4],
    3: [1],
    4: [2]
}
print("Graph:", graph)

"""
Min Heap
"""
import heapq

# Create a min-heap and push/pop elements.
heap = [5, 7, 9, 1, 3]
heapq.heapify(heap)               # Transform list into a heap.
print("Heap pop:", heapq.heappop(heap))  # Output: 1 (smallest element)

pq = []
heapq.heappush(pq, 1) # This works too

"""
Unicode 
"""
ord("A")

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


"""
Count
"""
arr = [1,2,3]
arr.count(2)


if __name__ == "__main__":
    cases = []

    for c in cases:
        print()
