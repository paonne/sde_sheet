## Rotate a Linked List

**❓**: Rotate the list to the right by k places.<br>
**🧠**:<br>
1. Find n % k where n is length of linked list, because that k is the correct value to rotate at the end<br>
2. Use slow and fast pointers. Give fast pointer k steps ahead start.<br>
3. Now move slow and fast pointer simultaneoulsy till fast pointer reaches end to find the kth point<br>
4. Change the corresponding node next values
=== "🐍"
    ```py
    --8<-- "07_linked_list_and_arrays/1_rotate_linked_list.py"
    ```
=== "🐋"
    ```go
    --8<-- "07_linked_list_and_arrays/1_rotate_linked_list.go"
    ```

[📘](https://takeuforward.org/data-structure/rotate-a-linked-list/) [💻](https://leetcode.com/problems/rotate-list/)<br>

---

## 3Sum

**❓**: Given an array of N integers, your task is to find unique triplets that add up to give a sum of zero. In short, you need to return an array of all the unique triplets [arr[a], arr[b], arr[c]] such that i!=j, j!=k, k!=i, and their sum is equal to zero.<br>
**Algorithm1 - Using HashMap**:<br>
1. Iterate through array. For every iteration, perform 2 sum using hashmap<br>
2. Sort the array to avoid duplicates<br>

**Algorithm2 - Using BinarySearch**:<br>
1. Sort the array<br>
2. Iterate through array. For every iteration, perform binary search approach<br>
3. Use left and right pointers. If sum is greater than zero, increase left else decrease right.<br>
4. If sum is zero, add items to res<br>
=== "🐍"
    ```py
    --8<-- "07_linked_list_and_arrays/3_3sum.py"
    ```
=== "🦀"
    ```rust
    --8<-- "07_linked_list_and_arrays/3_3sum.rs"
    ```
=== "🐋"
    ```go
    --8<-- "07_linked_list_and_arrays/3_3sum.go"
    ```
[📘](https://takeuforward.org/data-structure/3-sum-find-triplets-that-add-up-to-a-zero/) [💻](https://leetcode.com/problems/3sum/)<br>

---

## Trapping Rainwater

**❓**: Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

**Example**:  

![rainwater](https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png)

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]

Output: 6

Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

**🧠**:  
At a particular index, amount of water that can be trapped is height - min(prefix_max_height, suffix_max_height)
=== "🐍"
    ```py
    --8<-- "07_linked_list_and_arrays/4_trapping_rainwater.py"
    ```
=== "🦀"
    ```rust
    --8<-- "07_linked_list_and_arrays/4_trapping_rainwater.rs"
    ```
=== "🐋"
    ```go
    --8<-- "07_linked_list_and_arrays/4_trapping_rainwater.go"
    ```
[📘](https://takeuforward.org/data-structure/trapping-rainwater/) [💻](https://leetcode.com/problems/trapping-rain-water/description/)<br>

---

## Remove duplicates from sorted array

**❓**: Given an integer array sorted in non-decreasing order, remove the duplicates in place such that each unique element appears only once. The relative order of the elements should be kept the same.<br>
**🧠**:<br>
1. Use a curr pointer. Till curr pointer elements are correct<br>
2. Iterate through array and find whether current elements is same as previous. If not, put that value in curr pointer and increase curr pointer value<br>
3. The required length is till curr<br>
=== "🐍"
    ```py
    --8<-- "07_linked_list_and_arrays/5_remove_duplicates_from_sorted_array.py"
    ```
=== "🦀"
    ```rust
    --8<-- "07_linked_list_and_arrays/5_remove_duplicates_from_sorted_array.rs"
    ```
=== "🐋"
    ```go
    --8<-- "07_linked_list_and_arrays/5_remove_duplicates_from_sorted_array.go"
    ```
[📘](https://takeuforward.org/data-structure/remove-duplicates-in-place-from-sorted-array/) [💻](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)<br>

---

## Max Consecutive Ones

**❓**: Given an array that contains only 1 and 0 return the count of maximum consecutive ones in the array.<br>
**🧠**:<br>
1. Iterate through array, it element is 1, increment counter<br>
2. Else, res is maximum of res and curr<br>
=== "🐍"
    ```py
    --8<-- "07_linked_list_and_arrays/6_max_consecutive_ones.py"
    ```
=== "🦀"
    ```rust
    --8<-- "07_linked_list_and_arrays/6_max_consecutive_ones.rs"
    ```
=== "🐋"
    ```go
    --8<-- "07_linked_list_and_arrays/6_max_consecutive_ones.go"
    ```
[📘](https://takeuforward.org/data-structure/count-maximum-consecutive-ones-in-the-array/) [💻](https://leetcode.com/problems/max-consecutive-ones/)<br>