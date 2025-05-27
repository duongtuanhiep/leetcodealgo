"""
An insane guide on BS:
https://leetcode.com/problems/ugly-number-iii/solutions/769707/python-clear-explanation-powerful-ultimate-binary-search-template-solved-many-problems

As for the question "When can we use binary search?", my answer is that, If we can discover some kind of monotonicity, for example, if condition(k) is True then condition(k + 1) is True, then we can consider binary search.

"""


def binary_search(array) -> int:
    def condition(value) -> bool:
        pass

    left, right = 0, len(array)
    while left < right:
        mid = left + (right - left) // 2
        if condition(mid):
            right = mid
        else:
            left = mid + 1
    return left
