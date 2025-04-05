
"""
Removing largest digit first. 
so basically, 0 -> 9 AND if

Greedy algo: 
- Try to remove the largest digits first from left side since they'll show up first.
- 

"""

class Solution:
    def removeKdigits(self, num: str, k: int) -> str:
        stack = []
        
        # Process each digit
        for digit in num:
            # While there are still removals left and we can remove a larger digit from the stack:
            while k > 0 and stack and stack[-1] > digit:
                stack.pop()
                k -= 1
            stack.append(digit)
            
        # If k > 0, remove the last k digits (they are the largest remaining ones)
        if k:
            stack = stack[:-k]
        
        # Build the final number string and remove leading zeros
        result = ''.join(stack).lstrip('0')
        
        # Return '0' if the result is empty (i.e., the number was reduced to nothing)
        return result if result else "0"
