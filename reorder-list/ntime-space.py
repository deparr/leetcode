# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        stack = []
        curr = head
        while curr:
            stack.append(curr)
            curr = curr.next
        
        sz = len(stack)

        curr = head
        for _ in range(sz // 2):
            tmp = curr.next
            end = stack.pop()
            curr.next = end
            end.next = tmp
            curr = tmp

        curr.next = None
