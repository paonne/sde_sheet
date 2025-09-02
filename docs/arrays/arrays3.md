## Search in 2d Matrix

**❓**: Search in a sorted 2D matrix<br>

**Approach 1**:<br>
Simple Maths Equations<br>
1. Iterate through each row and check whether target can exist in that row or not<br>
2. If yes, do a binary search<br>
3. From those 2 equations, X and Y can be found out<br>

Time Complexity -> O(n + logm)
=== "🐍"
    ```py
    --8<-- "03_arrays_part3/1_search_in_2d_matrix.py"
    ```
=== "🦀"
    ```rust
    --8<-- "03_arrays_part3/1_search_in_2d_matrix.rs"
    ```
=== "🐋"
    ```go
    --8<-- "03_arrays_part3/1_search_in_2d_matrix.go"
    ```
[📘](https://takeuforward.org/data-structure/search-in-a-sorted-2d-matrix/)[💻](https://leetcode.com/problems/search-a-2d-matrix/)<br>

---

## Pow(x,n)

**❓**: Given a double x and integer n, calculate x raised to power n. Basically Implement pow(x, n).<br>

**🧠**<br>
```
6 ** 5

ans = 6
6 ** 4

(6 * 6) ** 2
36 ** 2
(36**36) ** 1
ans = ans * 36*36

1. If power if odd, multiply the curr number to ans and reduce curr number by 1
2. If power if even, square the curr number and reduce the power by half
```


TC - O(logn)<br>
SC - O(1)
=== "🐍"
    ```py
    --8<-- "03_arrays_part3/2_pow_x_of_n.py"
    ```
=== "🦀"
    ```rust
    --8<-- "03_arrays_part3/2_pow_x_of_n.rs"
    ```
=== "🐋"
    ```go
    --8<-- "03_arrays_part3/2_pow_x_of_n.go"
    ```
[📘](https://takeuforward.org/data-structure/implement-powxn-x-raised-to-the-power-n/)[💻](https://leetcode.com/problems/powx-n/)<br>

---

## Majority Element (>n/2 times)

**❓**: Given an array of N integers, write a program to return an element that occurs more than N/2 times in the given array. You may consider that such an element always exists in the array.<br>
=== "🐍"
    ```py
    --8<-- "03_arrays_part3/3_majority_element.py"
    ```
=== "🦀"
    ```rust
    --8<-- "03_arrays_part3/3_majority_element.rs"
    ```
=== "🐋"
    ```go
    --8<-- "03_arrays_part3/3_majority_element.go"
    ```
[📘](ttps://takeuforward.org/data-structure/find-the-majority-element-that-occurs-more-than-n-2-times/) [💻](https://leetcode.com/problems/majority-element/)<br>



---

## Majority Element (>n/3 times)

**❓**: Given an array of N integers. Find the elements that appear more than N/3 times in the array. If no such element exists, return an empty vector.<br>



**🧠**:<br>
Same as previous majority element algo. The difference is we apply the same algo for 2 elements and at the end we will find out which occurs more than n/3 times manually again.<br>

![image](https://static.takeuforward.org/wp/uploads/2023/04/Screenshot-2023-04-20-224857.png)

Time - O(2n)<br>
Space - O(1)
=== "🐍"
    ```py
    --8<-- "03_arrays_part3/4_majority_element_2.py"
    ```
=== "🦀"
    ```rust
    --8<-- "03_arrays_part3/4_majority_element_2.rs"
    ```
=== "🐋"
    ```go
    --8<-- "03_arrays_part3/4_majority_element_2.go"
    ```
[📘](https://takeuforward.org/data-structure/majority-elementsn-3-times-find-the-elements-that-appears-more-than-n-3-times-in-the-array/) [💻](https://leetcode.com/problems/majority-element-ii/)<br>

---

## Grid Unique Paths

**❓**: Given a matrix m X n, count paths from left-top to the right bottom of a matrix with the constraints that from each cell you can either only move to the rightward direction or the downward direction.<br>

![image](https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png)
=== "🐍"
    ```py
    --8<-- "03_arrays_part3/5_grid_unique_paths.py"
    ```
=== "🦀"
    ```rust
    --8<-- "03_arrays_part3/5_grid_unique_paths.rs"
    ```
=== "🐋"
    ```go
    --8<-- "03_arrays_part3/5_grid_unique_paths.go"
    ```
[📘](https://takeuforward.org/data-structure/grid-unique-paths-count-paths-from-left-top-to-the-right-bottom-of-a-matrix/) [💻](https://leetcode.com/problems/unique-paths/)<br>


---

## Reverse Pairs

Coming Soon