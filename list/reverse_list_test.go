package list

import (
    "testing"
)

// 反转单链表
// https://leetcode-cn.com/problems/reverse-linked-list/description/  easy

// 25. K 个一组翻转链表
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

// 92. 反转链表 II
// https://leetcode-cn.com/problems/reverse-linked-list-ii/
func Test_ReverseList(t *testing.T) {
    p := NewList()
    //for p != nil {
    //	t.Log(p)
    //	p = p.Next
    //}
    r := reverseListByArr(p)
    for r != nil {
        t.Log(r.Val)
        r = r.Next
    }
}

// 将链表结点存入数组，反向遍历数组得到反转链表 时间复杂度2N，空间复杂度2N
// 通过双指针遍历链表，不断反转next到pre指针，完成遍历
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return head
    }
    var pre, tmp *ListNode
    current := head
    for current != nil {
        tmp = current.Next
        current.Next = pre
        pre = current
        current = tmp
    }
    return pre
}

func reverseListByArr(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    arr := make([]*ListNode, 0)
    for head != nil {
        arr = append(arr, head)
        head = head.Next
    }
    tmp := new(ListNode)
    r := tmp
    for i := len(arr) - 1; i >= 0; i-- {
        tmp.Next = arr[i]
        tmp.Next.Next = nil
        tmp = tmp.Next
    }
    return r.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    stack := make([]*ListNode, 0)
    hair := &ListNode{
        Next: head,
    }
    current := hair.Next
    tmp := hair
    count := 0

    for current != nil {
        stack = append(stack, current)
        count++
        current = current.Next
        if count == k {
            count = 0
            for len(stack) > 0 {
                index := len(stack) - 1
                stack[index].Next = nil
                tmp.Next = stack[index]
                stack = stack[:index]
                tmp = tmp.Next
            }
        }
    }
    if len(stack) > 0 {
        tmp.Next = stack[0]
    }
    return hair.Next
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if m >= n {
        return head
    }
    hair := &ListNode{Next: head}
    cur := hair
    cnt := 0
    preBegin, begin, end, sufEnd := new(ListNode), new(ListNode), new(ListNode), new(ListNode)
    for cur != nil {
        if cnt == m-1 {
            preBegin = cur
        }
        if cnt == m {
            begin = cur
            preBegin.Next = nil
        }
        if cnt == n {
            end = cur
			sufEnd = cur.Next
            end.Next = nil
            break
        }
		cnt++
        cur = cur.Next
    }
    // 翻转begin-end
    preBegin.Next = reverse(begin)
    cur = preBegin
    for cur.Next != nil {
    	cur = cur.Next
	}
	cur.Next = sufEnd
    return hair.Next
}

func reverse(begin *ListNode) *ListNode {
    if begin == nil {
        return nil
    }
    var pre, tmp *ListNode
    current := begin
    for current != nil {
        tmp = current.Next
        current.Next = pre
        pre = current
        current = tmp
    }
    return pre
}
