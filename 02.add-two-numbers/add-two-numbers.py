from __future__ import annotations


class ListNode(object):
    def __init__(self, value: int(0), next: ListNode = None):
        self.value = value
        self.next = next


def generate_list(arr) -> ListNode:
    head = None
    current = None
    for v in arr:
        p = ListNode(value=v, next=None)
        if head is None:
            head = p
        if current is None:
            current = p
        else:
            current.next = p
            current = current.next

    return head


def print_list(head: ListNode):
    while head != None:
        print(head.value, end='')
        if head.next != None:
            print(" ---> ", end='')
        head = head.next
    print("")


def calculate_value_and_carry(one: ListNode, two: ListNode, carry: int) -> tuple[int, int]:
    one_value = 0
    if one is not None:
        one_value = one.value
    two_value = 0
    if one is not None:
        two_value = two.value
    return int((one_value + two_value + carry) % 10), int((one_value + two_value + carry) / 10)


def add_numbers(one: ListNode, two: ListNode) -> ListNode:
    reserved = ListNode
    carry = 0

    current = reserved
    while (one is not None or two is not None or carry != 0):
        value, carry = calculate_value_and_carry(one, two, carry)
        current.next = ListNode(value=value, next=None)
        current = current.next
        if one is not None:
            one = one.next
        if two is not None:
            two = two.next
    return reserved.next


if __name__ == '__main__':
    one = generate_list([1, 3, 4, 9])
    two = generate_list([6, 8, 9, 9])
    print_list(one)
    print_list(two)
    print_list(add_numbers(one, two))
