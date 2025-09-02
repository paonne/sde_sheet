## Rotate Matrix

**❓**: Given a matrix, your task is to rotate the matrix 90 degrees clockwise.<br>

=== "🐍"
    ```py
    --8<-- "02_arrays_part2/1_rotate_matrix.py"
    ```
=== "🦀"
    ```rust
    --8<-- "02_arrays_part2/1_rotate_matrix.rs"
    ```
=== "🐋"
    ```go
    --8<-- "02_arrays_part2/1_rotate_matrix.go"
    ```
[📘](https://takeuforward.org/data-structure/rotate-image-by-90-degree/) [💻](https://leetcode.com/problems/rotate-image/)<br>

---

## Merge Overlapping Sub Intervals

**❓**: Given an array of intervals, merge all the overlapping intervals and return an array of non-overlapping intervals.<br>
=== "🐍"
    ```py
    --8<-- "02_arrays_part2/2_merge_overlapping_subintervals.py"
    ```
=== "🦀"
    ```rust
    --8<-- "02_arrays_part2/2_merge_overlapping_subintervals.rs"
    ```
=== "🐋"
    ```go
    --8<-- "02_arrays_part2/2_merge_overlapping_subintervals.go"
    ```
[📘](https://takeuforward.org/data-structure/merge-overlapping-sub-intervals/) [💻](https://leetcode.com/problems/merge-intervals/)<br>

---

## Merge 2 sorted arrays without extra space

**❓**: Given two sorted arrays arr1[] and arr2[] of sizes n and m in non-decreasing order. Merge them in sorted order. Modify arr1 so that it contains the first N elements and modify arr2 so that it contains the last M elements.<br>
=== "🐍"
    ```py
    --8<-- "02_arrays_part2/3_merge_two_sorted_arrays.py"
    ```
=== "🦀"
    ```rust
    --8<-- "02_arrays_part2/3_merge_two_sorted_arrays.rs"
    ```
=== "🐋"
    ```go
    --8<-- "02_arrays_part2/3_merge_two_sorted_arrays.go"
    ```
[📘](https://takeuforward.org/data-structure/merge-two-sorted-arrays-without-extra-space/) [💻](https://leetcode.com/problems/merge-sorted-array/)<br>

---

## Duplicate in array of n+1 integers

**❓**: Given an array of N + 1 size, where each element is between 1 and N. Assuming there is only one duplicate number, your task is to find the duplicate number.<br>

**🧠**:<br>
Floyd's Cycle Detection<br>
1. Slow moves one point at a time<br>
2. Fast moves 2 points at a time.<br>
3. After they colloide, if we reset first pointer, and move that and slow pointer at a time, they will colloide at our result.<br>

Please check  [here](https://www.youtube.com/watch?v=wjYnzkAhcNk) for formula and intuition<br>
=== "🐍"
    ```py
    --8<-- "02_arrays_part2/4_duplicate_of_n+1_integers.py"
    ```
=== "🦀"
    ```rust
    --8<-- "02_arrays_part2/4_duplicate_of_n+1_integers.rs"
    ```
=== "🐋"
    ```go
    --8<-- "02_arrays_part2/4_duplicate_of_n+1_integers.go"
    ```
[📘](https://takeuforward.org/data-structure/find-the-duplicate-in-an-array-of-n1-integers/) [💻](https://leetcode.com/problems/find-the-duplicate-number/)<br>

---

## Repeating and Missing number

**❓**: You are given a read-only array of N integers with values also in the range [1, N] both inclusive. Each integer appears exactly once except A which appears twice and B which is missing. The task is to find the repeating and missing numbers A and B where A repeats twice and B is missing.<br>

**🧠**:<br>
Simple Maths Equations<br>
1. Find sum of n numbers and find X - Y value. (X- Repeating, Y - Missing)<br>
2. Find sum of squares of n numbers and form second equation. (X2 - Y2 = (X+Y)(X-Y)). From which we can get X + Y value<br>
3. From those 2 equations, X and Y can be found out<br>

Please check [here](https://www.youtube.com/watch?v=wjYnzkAhcNk) for formula and intuition<br>
=== "🐍"
    ```py
    --8<-- "02_arrays_part2/5_repeating_and_missing_number.py"
    ```
=== "🦀"
    ```rust
    --8<-- "02_arrays_part2/5_repeating_and_missing_number.rs"
    ```
=== "🐋"
    ```go
    --8<-- "02_arrays_part2/5_repeating_and_missing_number.go"
    ```
[📘](https://takeuforward.org/data-structure/find-the-repeating-and-missing-numbers/)<br>

---

## Inversion of Array

Coming Soon

