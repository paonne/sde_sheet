## 2 Sum
**❓**: Check if a pair with given sum exists in Array<br>
**🧠**:<br>
1. Store the current element with index in hash map<br>
2. while iterating through arrays, check whether curr - sum in hashmap<br>

=== "🐍"
    ```py
    --8<-- "04_arrays_part4/1_2sum.py"
    ```
=== "🦀"
    ```rust
    --8<-- "04_arrays_part4/1_2sum.rs"
    ```
=== "🐋"
    ```go
    --8<-- "04_arrays_part4/1_2sum.go"
    ```
[📘](https://takeuforward.org/data-structure/two-sum-check-if-a-pair-with-given-sum-exists-in-array/) [💻](https://leetcode.com/problems/two-sum/)<br>

---

## 4 Sum
**❓**: Find Quads that add up to a target value<br>
**🧠**:<br>
1. Sort the array and use 4 pointers.<br>
2. Fix i and j pointers and put left and right pointers at the left and right end of remaining array<br>
3. **To avoid duplicate quads, while moving the pointers, move them only if they are not equal to their previous elements**<br>
4. If 4 elements sum is equal to target, store them. Else according to sum, move left and right pointers as elements are already sorted


```py
{!04_arrays_part4/2_4sum.py!lines=1-28}
```

(Another approach to avoid duplicates is storing them in set)
=== "🐍"
    ```py
    --8<-- "04_arrays_part4/2_4sum.py:30:48"
    ```
=== "🦀"
    ```rust
    --8<-- "04_arrays_part4/2_4sum.rs"
    ```
=== "🐋"
    ```go
    --8<-- "04_arrays_part4/2_4sum.go"
    ```

[📘](https://takeuforward.org/data-structure/4-sum-find-quads-that-add-up-to-a-target-value/) [💻](https://leetcode.com/problems/4sum/)<br>


---

## Longest Consecutive Sequence
**❓**: You are given an array of ‘N’ integers. You need to find the length of the longest sequence which contains the consecutive elements.<br>
**🧠**:<br>
O(NlogN) Solution<br>
1. Sort the array<br>
2. Find the consecutive sequence by iterating through array. (it cannot have same elements)<br>

O(N) Solution<br>
1. Store all the elements in set<br>
2. For every element, search and count all the consecutive elements only if<br>
    a. It is start of sequence(curr-1 is not in set)<br>
    b. We will storing maximum sequence in a variable res. if curr + res is not present in set, that means the curr count itself is large sequence, again don't need to find the current sequence length
=== "🐍"
    ```py
    --8<-- "04_arrays_part4/3_longest_consecutive_sequence.py"
    ```
=== "🦀"
    ```rust
    --8<-- "04_arrays_part4/3_longest_consecutive_sequence.rs"
    ```
=== "🐋"
    ```go
    --8<-- "04_arrays_part4/3_longest_consecutive_sequence.go"
    ```
[📘](https://takeuforward.org/data-structure/longest-consecutive-sequence-in-an-array/) [💻](https://leetcode.com/problems/longest-consecutive-sequence/)<br>

---

## Longest SubArray with Zero Sum
**❓**: Given an array containing both positive and negative integers, we have to find the length of the longest subarray with the sum of all elements equal to zero.<br>
**🧠**:<br>
1. Store the prefix sum in hashmap with value as index<br>
2. if we find the prefix sum again, then that part of array sum is zero<br>
=== "🐍"
    ```py
    --8<-- "04_arrays_part4/4_longest_subarray_with_sum_zero.py"
    ```
=== "🦀"
    ```rust
    --8<-- "04_arrays_part4/4_longest_subarray_with_sum_zero.rs"
    ```
=== "🐋"
    ```go
    --8<-- "04_arrays_part4/4_longest_subarray_with_sum_zero.go"
    ```
[📘](https://takeuforward.org/data-structure/length-of-the-longest-subarray-with-zero-sum/)<br>

---

## Count number of subarrays with given XOR K

Coming Soon


---

## Longest Substring without repeating characters 
**❓**: Given a String, find the length of longest substring without any repeating character.<br>
**🧠**:<br>
1. Store the seen elements in hashmap with its indexes<br>
2. Once you encounter a seen element, compute the res by finding the gap b/w current index and seen index and reset left pointer to seen index plus 1<br>
=== "🐍"
    ```py
    --8<-- "04_arrays_part4/6_longest_substring_without_repeating_character.py"
    ```
=== "🦀"
    ```rust
    --8<-- "04_arrays_part4/6_longest_substring_without_repeating_character.rs"
    ```
=== "🐋"
    ```go
    --8<-- "04_arrays_part4/6_longest_substring_without_repeating_character.go"
    ```
[📘](https://takeuforward.org/data-structure/length-of-the-longest-subarray-with-zero-sum/) [💻](https://leetcode.com/problems/longest-substring-without-repeating-characters/)<br>

