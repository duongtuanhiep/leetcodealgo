"""

Question 2570: https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/

"""


class Solution:
    def mergeArrays(
        self, nums1: List[List[int]], nums2: List[List[int]]
    ) -> List[List[int]]:
        # Merging two sorted arrays so that
        i, j, res = 0, 0, []
        while i < len(nums1) and j < len(nums2):
            if nums1[i][0] < nums2[j][0]:
                res.append(nums1[i])
                i += 1
            elif nums1[i][0] > nums2[j][0]:
                res.append(nums2[j])
                j += 1
            else:
                res.append([nums1[i][0], nums1[i][1] + nums2[j][1]])
                i += 1
                j += 1

        res.extend(nums1[i:] or nums2[j:])  # or res += nums1[i:] \n res += nums2[j:]
        return res
