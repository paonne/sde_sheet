## Reverse a Linked List

**❓**: Given the head of a singly linked list, write a program to reverse the linked list, and return the head pointer to the reversed list.<br>

=== "🐍"
    ```py
    --8<-- "04_linked_list/1_reverse_a_linked_list.py"
    ```
=== "🐋"
    ```go
    --8<-- "04_linked_list/1_reverse_a_linked_list.go"
    ```

[📘](https://takeuforward.org/data-structure/reverse-a-linked-list/) [💻](https://leetcode.com/problems/reverse-linked-list/)<br>

---

## Middle of a Linked List

**❓**: Given the head of a linked list of integers, determine the middle node of the linked list. However, if the linked list has an even number of nodes, return the second middle node.<br>
**🧠**:<br>
1. Use slow and fast pointers where fast pointer moves 2 steps and slow pointers moves 1 step at a time<br>
2. When fast pointer is at the end of linked list, slow pointer will be in the middle<br>

=== "🐍"
    ```py
    --8<-- "04_linked_list/2_middle_of_linked_list.py"
    ```
=== "🐋"
    ```go
    --8<-- "04_linked_list/2_middle_of_linked_list.go"
    ```

[📘](https://takeuforward.org/data-structure/find-middle-element-in-a-linked-list/) [💻](https://leetcode.com/problems/middle-of-the-linked-list/description/)<br>

---

## Remove nth node from end

**❓**: Given a linked list and an integer N, the task is to delete the Nth node from the end of the linked list and print the updated linked list.<br>
**🧠**:<br>
1. Use 2 pointers slow and fast. Move fast pointer first from head by n steps<br>
2. Now move the slow pointer till fast pointer reaches end.<br>
3. Now the slow pointer's next node needs to be removed<br>
4. Handle the edge case in the start if fast pointer moves beyond linked list<br>

=== "🐍"
    ```py
    --8<-- "04_linked_list/4_remove_nth_node_from_end.py"
    ```
=== "🐋"
    ```go
    --8<-- "04_linked_list/4_remove_nth_node_from_end.go"
    ```

[📘](https://takeuforward.org/data-structure/remove-n-th-node-from-the-end-of-a-linked-list/) [💻](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)<br>

---

## Add 2 Numbers

**❓**: Given the heads of two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.<br>
**🧠**:<br>
1. Iterate through both linked lists by carrying current remainder<br>
2. Create new node with the remainder and curr sum and return the final linked list<br>

=== "🐍"
    ```py
    --8<-- "04_linked_list/5_add_2_numbers.py"
    ```
=== "🐋"
    ```go
    --8<-- "04_linked_list/5_add_2_numbers.go"
    ```

[📘](https://takeuforward.org/data-structure/add-two-numbers-represented-as-linked-lists/) [💻](https://leetcode.com/problems/add-two-numbers/)<br>

---

## Delete the given node

**❓**: Write a function to delete a node in a singly-linked list. You will not be given access to the head of the list instead, you will be given access to the node to be deleted directly. It is guaranteed that the node to be deleted is not a tail node in the list.<br>
**🧠**:<br>
Replace the node value with next node value and make next node as next node's next<br>

=== "🐍"
    ```py
    --8<-- "04_linked_list/6_delete_the_given_node.py"
    ```
=== "🐋"
    ```go
    --8<-- "04_linked_list/6_delete_the_given_node.go"
    ```

[📘](https://takeuforward.org/data-structure/delete-given-node-in-a-linked-list-o1-approach/) [💻](https://leetcode.com/problems/delete-node-in-a-linked-list/)<br>