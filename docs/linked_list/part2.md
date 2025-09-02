## Find the intersection of 2 linked lists

**❓**: Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return null.<br>
**🧠**:<br>
1. First find the length difference between 2 linked lists<br>
2. Use 2 pointers cur1, curr2 for 1st and 2nd linked lists respectively<br>
3. If l1 > l2, move curr1 by l1-l2 steps ahead, else do the opposite for curr2<br>
4. Now, curr1 and curr2 are same distance apart from intersection, move them simultaneously to find the junction point<br>
=== "🐍"
    ```py
    --8<-- "06_linked_list_2/1_find_intersection_of_linked_lists.py"
    ```
=== "🐋"
    ```go
    --8<-- "06_linked_list_2/1_find_intersection_of_linked_lists.go"
    ```

[📘](https://takeuforward.org/data-structure/find-intersection-of-two-linked-lists/) [💻](https://leetcode.com/problems/intersection-of-two-linked-lists/)<br>

---

## Detect a cycle in Linked List

**🧠**:<br>
1. Use slow and fast pointers where fast moves 2 steps at a time and slow moves 1 step at a time<br>
2. If there is a cycle, they will meet each other<br>
=== "🐍"
    ```py
    --8<-- "06_linked_list_2/2_detect_a_cycle_in_linked_list.py"
    ```
=== "🐋"
    ```go
    --8<-- "06_linked_list_2/2_detect_a_cycle_in_linked_list.go"
    ```

[📘](https://takeuforward.org/data-structure/detect-a-cycle-in-a-linked-list/) [💻](https://leetcode.com/problems/linked-list-cycle/description/)<br>

---

## Check for Palindrome

**❓**: Check if the given Linked List is Palindrome<br>

**🧠**:<br>
1. Use slow and fast pointers and find the mid of the linked list<br>
2. Reverse the 2nd half of linked list<br>
3. Now travese both nodes and check values are same or not<br>

=== "🐍"
    ```py
    --8<-- "06_linked_list_2/4_check_for_palindrome.py"
    ```
=== "🐋"
    ```go
    --8<-- "06_linked_list_2/4_check_for_palindrome.go"
    ```

[📘](https://takeuforward.org/data-structure/check-if-given-linked-list-is-plaindrome/) [💻](https://leetcode.com/problems/palindrome-linked-list/)<br>

---

## Starting Point of Loop

**❓**: Given the head of a linked list that may contain a cycle, return the starting point of that cycle. If there is no cycle in the linked list return null.<br>

**🧠**:<br>
1. Use slow and fast pointers and and find whether loop is there or not<br>
2. #TODO If yes, now the distance between required node, head and required node and slow is same(Didn't get the approach completely)<br>

=== "🐍"
    ```py
    --8<-- "06_linked_list_2/5_starting_point_of_loop.py"
    ```
=== "🐋"
    ```go
    --8<-- "06_linked_list_2/5_starting_point_of_loop.go"
    ```

[📘](https://takeuforward.org/data-structure/starting-point-of-loop-in-a-linked-list/) [💻](https://leetcode.com/problems/linked-list-cycle-ii/)<br>