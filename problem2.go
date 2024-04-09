/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var newNode *ListNode
    var a []int
    var b []int
    var c []int

    for l1.Next != nil {
        a = append(a,l1.Val)
        l1 = l1.Next
    }

    a = append(a,l1.Val)

    for l2.Next != nil {
        b = append(b,l2.Val)
        l2 = l2.Next
    }

    b = append(b,l2.Val)

    var tempa []int
    var tempb []int

    for i := len(a) - 1; i >= 0; i-- {
        tempa = append(tempa, a[i])
    }

    for i := len(b) - 1; i >= 0; i-- {
        tempb = append(tempb, b[i])
    }

    if len(tempa) > len(tempb) {
        for i := len(tempb); i <= len(tempa) ; i++ {
            if i != (len(a)){
                tempb = append([]int{0}, tempb...)
            }
        }
    } else if len(tempa) < len(tempb) {
        for i := len(tempa); i < len(tempb); i++ {
            if i != (len(tempb)) {
                tempa = append([]int{0}, tempa...)
            }
        }
    }

    for i := len(tempa) - 1; i >= 0; i-- {
        c = append(c, tempa[i] + tempb[i])
    }

    
    temp := 0
    for j := 0; j < len(c); j++ {
        c[j] = c[j] + temp;
        if c[j] >= 10 {
            c[j] = c[j] % 10
            temp = 1
        } else {
            temp = 0
        }

        if j == len(c) - 1{
            if c[j] >= 10 && temp == 0{
                c[j] = c[j] % 10
                c = append(c, c[j])
            } else if temp == 1 {
                c = append(c, 1)
            }
            break
        }
    }

	var head *ListNode
	var tail *ListNode 

	for _, val := range c {
		newNode = &ListNode{Val: val} 
		if head == nil {
			head = newNode
		} else {
			tail.Next = newNode
		}
		tail = newNode
	}

    return head
}
