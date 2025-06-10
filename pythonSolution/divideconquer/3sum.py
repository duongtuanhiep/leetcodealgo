from typing import List

"""

Question 15: https://leetcode.com/problems/3sum/?envType=problem-list-v2&envId=27ol1gds

"""


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        # Sort the array
        nums.sort()
        result = []
        
        # Loop through the numbers
        for i in range(len(nums) - 2):
            # Skip duplicate values for the first element of the triplet
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            
            # Set two pointers for the remaining two numbers
            left, right = i + 1, len(nums) - 1
            
            while left < right:
                total = nums[i] + nums[left] + nums[right]
                
                # If the sum is zero, we found a valid triplet
                if total == 0:
                    result.append([nums[i], nums[left], nums[right]])
                    
                    # Skip duplicates for the left pointer
                    while left < right and nums[left] == nums[left + 1]:
                        left += 1
                    # Skip duplicates for the right pointer
                    while left < right and nums[right] == nums[right - 1]:
                        right -= 1
                    
                    # Move both pointers after storing a valid triplet
                    left += 1
                    right -= 1
                elif total < 0:
                    left += 1  # Need a larger sum
                else:
                    right -= 1  # Need a smaller sum
        return result

if __name__ == "__main__":
    cases = [
        [-1,0,1,2,-1,-4],
        [0,1,1],
        [0,0,0,0],
        [-4,-2,1,-5,-4,-4,4,-2,0,4,0,-2,3,1,-5,0]
    ]
    for c in cases: 
        print(Solution().threeSum(c))
