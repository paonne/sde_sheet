## Nth root of an integer

**❓**: You are given two positive integers 'n' and 'm'. You have to return the 'nth' root of 'm', i.e. 'm(1/n)'. If the 'nth root is not an integer, return -1. 

**Example**:  
3 27

Output:
3

Explanation For Sample Input 1:
3rd Root of 27 is 3, as (3)^3 equals 27.
=== "🐍"
    ```py
    --8<-- "11_binary_search/01_nth_root.py"
    ```
=== "🦀"
    ```rust
    --8<-- "11_binary_search/01_nth_root.rs"
    ```
=== "🐋"
    ```go
    --8<-- "11_binary_search/01_nth_root.go"
    ```
[📘](https://takeuforward.org/data-structure/nth-root-of-a-number-using-binary-search/) [💻](https://www.naukri.com/code360/problems/1062679?topList=striver-sde-sheet-problems&leftPanelTabValue=PROBLEM)<br>

---

## Matrix Median

Coming Soon

---

## Single Element in a Sorted Array

**❓**: You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once.

Return the single element that appears only once.

**Example**:  
Input: nums = [1,1,2,3,3,4,4,8,8]
Output: 2

**🧠**:  
Elements which are repeated twice will be at   
1. odd, even indices to left of target  
2. even, odd indices to right of target
=== "🐍"
    ```py
    --8<-- "11_binary_search/03_single_element_in_sorted_array.py"
    ```
=== "🦀"
    ```rust
    --8<-- "11_binary_search/03_single_element_in_sorted_array.rs"
    ```
=== "🐋"
    ```go
    --8<-- "11_binary_search/03_single_element_in_sorted_array.go"
    ```
[📘](https://takeuforward.org/data-structure/search-single-element-in-a-sorted-array/) [💻](https://leetcode.com/problems/single-element-in-a-sorted-array/description/)<br>

---

## Search in Rotated Sorted Array

**❓**: There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

**Example**:  
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

**🧠**:  
1. On doing binary search, atleast one of left or right half will be sorted.  
2. Eliminate one of them using that property.
=== "🐍"
    ```py
    --8<-- "11_binary_search/04_search_in_rotated_sorted_array.py"
    ```
=== "🦀"
    ```rust
    --8<-- "11_binary_search/04_search_in_rotated_sorted_array.rs"
    ```
=== "🐋"
    ```go
    --8<-- "11_binary_search/04_search_in_rotated_sorted_array.go"
    ```
[📘](https://takeuforward.org/data-structure/search-element-in-a-rotated-sorted-array/) [💻](https://leetcode.com/problems/search-in-rotated-sorted-array/description/)<br>

---

## Median of 2 Sorted Arrays

Coming Soon

---

## Kth Element of 2 Sorted Arrays

Coming Soon

---

## Allocate Minimum Number of Pages

Coming Soon

---

## Agressive Cows

**❓**: Given an array nums of size n, which denotes the positions of stalls, and an integer k, which denotes the number of aggressive cows, assign stalls to k cows such that the minimum distance between any two cows is the maximum possible. Find the maximum possible minimum distance.    

**Example**:  
Input: n = 6, k = 4, nums = [0, 3, 4, 7, 10, 9]

Output: 3

Explanation: The maximum possible minimum distance between any two cows will be 3 when 4 cows are placed at positions [0, 3, 7, 10]. Here the distances between cows are 3, 4, and 3 respectively. We cannot make the minimum distance greater than 3 in any ways.  

**🧠**:  
1. Minimum possible way is 1 and maximum is max(arr) - min(arr).  
2. Linear search approach would be trying all the values in range till we find maximum result.  
3. Binary search approach is going to mid element and check whether it is possible or not.  
4. If it is possible, as we need max value, eliminate left half to find better answer in right half.  
5. Else, eliminate right half to find answer in left half.

```py
{!11_binary_search/08_agressive_cows.py!}
```

[📘](https://takeuforward.org/data-structure/aggressive-cows-detailed-solution/) <br>

---