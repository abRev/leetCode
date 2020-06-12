func getIntersectionNode(headA, headB *ListNode) *ListNode {
    left:=0
    right:=0
    for head:=headA;head!=nil;head = head.Next {
        left++
    }
    for head:=headB;head!=nil;head = head.Next {
        right++
    }
    for left > right {
        headA = headA.Next
        left--
    }
    for left < right {
        headB = headB.Next
        right--
    }
    for headA !=nil {
        if headA == headB {
            return headA
        }
        headA = headA.Next
        headB = headB.Next
    }
    return nil
}