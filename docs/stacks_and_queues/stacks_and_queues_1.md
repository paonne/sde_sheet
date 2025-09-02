## Check for Balanced Parentheses

**❓**: Check Balanced Parentheses. Given string str containing just the characters '(', ')', '{', '}', '[' and ']', check if the input string is valid and return true if the string is balanced otherwise return false.   

**Example**:  
Input: str = “( )[ { } ( ) ]”

Output: True

Explanation: As every open bracket has its corresponding 
close bracket. Match parentheses are in correct order 
hence they are balanced.

**🧠**:  
1. Iterate through string, if opened brace is found append it to the stack.  
2. If close brace if found, stack top element should be corresponding open brace.  
3. If stack is empty or top element is a mismatch, return False.  
4. Else pop the top element and continue.  
5. At the end of iteration, stack should be empty.  
=== "🐍"
    ```py
    --8<-- "13_stacks_and_queues/balanced_paranthesis.py"
    ```
=== "🦀"
    ```rust
    --8<-- "13_stacks_and_queues/balanced_paranthesis.rs"
    ```
=== "🐋"
    ```go
    --8<-- "13_stacks_and_queues/balanced_paranthesis.go"
    ```

[📘](https://takeuforward.org/data-structure/check-for-balanced-parentheses/) [💻](https://leetcode.com/problems/valid-parentheses/)<br>

---

## Next Greater Element Using Stack

**❓**: The next greater element of some element x in an array is the first greater element that is to the right of x in the same array.

You are given two distinct 0-indexed integer arrays nums1 and nums2, where nums1 is a subset of nums2.

For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j] and determine the next greater element of nums2[j] in nums2. If there is no next greater element, then the answer for this query is -1.

Return an array ans of length nums1.length such that ans[i] is the next greater element as described above.

**Example**:  
Input: nums1 = [4,1,2], nums2 = [1,3,4,2]  
Output: [-1,3,-1]  
Explanation: The next greater element for each value of nums1 is as follows:  
- 4 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.  
- 1 is underlined in nums2 = [1,3,4,2]. The next greater element is 3.  
- 2 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.  

**🧠**:  
1. Use stack to find nge.  
2. Store nge as hash map because we need nge for corresponding elements from nums1.  
3. Return result array using nge map.
=== "🐍"
    ```py
    --8<-- "13_stacks_and_queues/next_greater_element.py"
    ```
=== "🦀"
    ```rust
    --8<-- "13_stacks_and_queues/next_greater_element.rs"
    ```
=== "🐋"
    ```go
    --8<-- "13_stacks_and_queues/next_greater_element.go"
    ```

[📘](https://takeuforward.org/data-structure/next-greater-element-using-stack/) [💻](https://leetcode.com/problems/next-greater-element-i/)<br>

---