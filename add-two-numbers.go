
//AddValue function
func AddValue(nodeA *ListNode, nodeB *ListNode, overflow int) {
	if nodeA == nil{
		return;
	}
	result := nodeA.Val + nodeB.Val + overflow

	nodeA.Val = result % 10
	c := 0

	if result >= 10{
		c = 1
	}

	if nodeA.Next != nil || nodeB.Next != nil{
		if nodeA.Next == nil{
			nodeA.Next = &ListNode{Val: 0}
		}
		if nodeB.Next == nil{
			nodeB.Next = &ListNode{Val: 0}
		}
		AddValue(nodeA.Next, nodeB.Next, c)
	
	} else if result >= 10{
		nodeA.Next = &ListNode{Val: c}
	}
    
    
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	AddValue(l1, l2, 0)
	return l1
     
}