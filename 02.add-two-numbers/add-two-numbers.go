package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

func generateList(numbers []int) *ListNode {
	var head *ListNode = nil
	var current *ListNode = nil

	for i := 0; i < len(numbers); i++ {
		p := &ListNode{
			value: numbers[i],
			next:  nil,
		}

		if current != nil {
			current.next = p
			current = current.next
		} else {
			current = p
		}

		if head == nil {
			head = p
		}
	}

	return head
}

func printList(head *ListNode) {

	for p := head; p != nil; p = p.next {
		fmt.Print(p.value)
		if p.next != nil {
			fmt.Print(" ---> ")
		}
	}
	fmt.Println()
}

func reversePrintList(head *ListNode) {
	if head.next != nil {
		reversePrintList(head.next)
		fmt.Print(" ---> ")
	}

	fmt.Print(head.value)
}

func addTwoNumbers(one *ListNode, two *ListNode) *ListNode {
	//return addTwoNumbersImpl(one, two, 0)
	reserved := &ListNode{}
	carry := 0
	value := 0

	current := reserved
	for one != nil || two != nil || carry != 0 {
		value, carry = calcValueCarry(one, two, carry)
		current.next = &ListNode{
			value: value,
			next:  nil,
		}
		current = current.next
		if one != nil {
			one = one.next
		}
		if two != nil {
			two = two.next
		}
	}

	return reserved.next
}

func calcValueCarry(one *ListNode, two *ListNode, carry int) (int, int) {
	val1 := 0
	if one != nil {
		val1 = one.value
	}

	val2 := 0
	if two != nil {
		val2 = two.value
	}

	return (val1 + val2 + carry) % 10, (val1 + val2 + carry) / 10
}

func addTwoNumbersImpl(one *ListNode, two *ListNode, carry int) *ListNode {
	if one == nil && two == nil && carry == 0 {
		return nil
	}

	value, newCarry := calcValueCarry(one, two, carry)
	var oneNext *ListNode = nil
	if one != nil {
		oneNext = one.next
	}

	var twoNext *ListNode = nil
	if two != nil {
		twoNext = two.next
	}

	return &ListNode{
		value: value,
		next:  addTwoNumbersImpl(oneNext, twoNext, newCarry),
	}
}

func main() {
	a := []int{2, 5, 8, 9}
	b := []int{7, 7, 2, 9}
	printList(generateList(a))
	printList(generateList(b))
	fmt.Println("print sum now: ")
	printList(addTwoNumbers(generateList(a), generateList(b)))

	reversePrintList(addTwoNumbers(generateList(a), generateList(b)))
}
