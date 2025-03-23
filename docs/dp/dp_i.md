## Max Product Subarray

**❓**: Given an array that contains both negative and positive integers, find the maximum product subarray.

**Example**:  
Input: [1,2,3,4,5,0]  
Output: 120  
Explanation: In the given array, we can see 1×2×3×4×5 gives maximum product value.  

Input: Nums = [1,2,-3,0,-4,-5]  
Output: 20  
Explanation: In the given array, we can see (-4)×(-5) gives maximum product value.  

=== "🐍"
    ```py
    --8<-- "25_dp/01_max_prod_subarray.py"
    ```
=== "🦀"
    ```rust
    --8<-- "25_dp/01_max_prod_subarray.rs"
    ```
    
[📘](https://takeuforward.org/data-structure/maximum-product-subarray-in-an-array/) [💻](https://leetcode.com/problems/maximum-product-subarray/)<br>

## Longest Increasing Subsequence

**❓**: Given an integer array nums, return the length of the longest strictly increasing subsequence.

**Example**:  
Input: [10,9,2,5,3,7,101,18]  
Output: 4   
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.  

Input: Nums = [0,1,0,3,2,3]   
Output: 4  

=== "🐍"
    ```py
    --8<-- "25_dp/02_longest_increasing_subsequence.py"
    ```
=== "🦀"
    ```rust
    --8<-- "25_dp/02_longest_increasing_subsequence.rs"
    ```
    
[📘](https://takeuforward.org/data-structure/longest-increasing-subsequence-dp-41/) [💻](https://leetcode.com/problems/longest-increasing-subsequence/)<br>

## Longest Common Subsequence

**❓**: Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

**Example**:  
Input: text1 = "abcde", text2 = "ace"  
Output: 3     
Explanation: The longest common subsequence is "ace" and its length is 3.   

Input: Nums = text1 = "abc", text2 = "abc"    
Output: 3  
Explanation: The longest common subsequence is "abc" and its length is 3.  

=== "🐍"
    ```py
    --8<-- "25_dp/03_longest_common_subsequence.py"
    ```
=== "🦀"
    ```rust
    --8<-- "25_dp/03_longest_common_subsequence.rs"
    ```
    
[📘](https://takeuforward.org/data-structure/longest-common-subsequence-dp-25/) [💻](https://leetcode.com/problems/longest-common-subsequence/)<br>

## 0/1 Knapsack

**❓**: A thief wants to rob a store. He is carrying a bag of capacity W. The store has ‘n’ items. Its weight is given by the ‘wt’ array and its value by the ‘val’ array. He can either include an item in its knapsack or exclude it but can’t partially have it as a fraction. We need to find the maximum value of items that the thief can steal.  

**Example**:  
![alternate](https://lh3.googleusercontent.com/bgxCoEMMmxwuN3HgXNgYZ_o_Gxb8QSyBXXTxwSpsP7757ECoAplkpCtBdoS5LFM-3C-YhAbIhIbB9XGKxTZBGm6GB7HstnwmKf3hix_8V3zfuHLrK70bOjr01PezXrYOymoEiCaZ) 

=== "🐍"
    ```py
    --8<-- "25_dp/04_01_knapsack.py"
    ```
=== "🦀"
    ```rust
    --8<-- "25_dp/04_01_knapsack.rs"
    ```
    
[📘](https://takeuforward.org/data-structure/0-1-knapsack-dp-19/) [💻](https://www.naukri.com/code360/problems/1072980)<br>

## Edit Distance

**❓**: We are given two strings ‘S1’ and ‘S2’. We need to convert S1 to S2. The following three operations are allowed:  
    1. Deletion of a character.  
    2. Replacement of a character with another one.  
    3. Insertion of a character.  
We have to return the minimum number of operations required to convert S1 to S2 as our answer.

![image](https://lh6.googleusercontent.com/Dxn0cvswqpu9nszd6gMXThvxbSwlyz_lLBUwzYmyNhvV9LcGNYWUjC9D8T9iP0pUlaf1WRtpYz061ttrSe8cvo-DvUeknkKX8MuDrBy4_JhsSqj4TVKoEoePOauIEpvN-UaeSZ5N)  

=== "🐍"
    ```py
    --8<-- "25_dp/05_edit_distance.py"
    ```
=== "🦀"
    ```rust
    --8<-- "25_dp/05_edit_distance.rs"
    ```
    
[📘](https://takeuforward.org/data-structure/edit-distance-dp-33/) [💻](https://leetcode.com/problems/edit-distance/)<br>

## Max Sum Increasing Subsequence

**❓**: Given an array of positive integers arr. Find the maximum sum subsequence of the given array such that the integers in the subsequence are sorted in strictly increasing order i.e. a strictly increasing subsequence. 

**Example**:  
Input: [1, 101, 2, 3, 100]    
Output: 106    
Explanation: The maximum sum of a increasing sequence is obtained from [1, 2, 3, 100].     

Input: arr[] = [4, 1, 2, 3]  
Output: 6  
Explanation: The maximum sum of a increasing sequence is obtained from {1, 2, 3}.  

=== "🐍"
    ```py
    --8<-- "25_dp/06_maximum_sum_increasing_subsequence.py"
    ```
=== "🦀"
    ```rust
    --8<-- "25_dp/06_maximum_sum_increasing_subsequence.rs"
    ```
    
[💻](https://www.geeksforgeeks.org/problems/maximum-sum-increasing-subsequence4749/1)<br>

## Matrix Chain Multiplication

**❓**: Given an array arr[] which represents the dimensions of a sequence of matrices where the ith matrix has the dimensions (arr[i-1] x arr[i]) for i>=1, find the most efficient way to multiply these matrices together. The efficient way is the one that involves the least number of multiplications.  

**Example**:  
Input: arr[] = [2, 1, 3, 4]  
Output: 20  
Explanation: There are 3 matrices of dimensions 2 × 1, 1 × 3, and 3 × 4, Let this 3 input matrices be M1, M2, and M3. There are two ways to multiply: ((M1 x M2) x M3) and (M1 x (M2 x M3)), note that the result of (M1 x M2) is a 2 x 3 matrix and result of (M2 x M3) is a 1 x 4 matrix. 
((M1 x M2) x M3)  requires (2 x 1 x 3) + (2 x 3 x 4) = 30 
(M1 x (M2 x M3))  requires (1 x 3 x 4) + (2 x 1 x 4) = 20. 
The minimum of these two is 20.   

=== "🐍"
    ```py
    --8<-- "25_dp/07_matrix_chain_multiplication.py"
    ```
=== "🦀"
    ```rust
    --8<-- "25_dp/07_matrix_chain_multiplication.rs"
    ```
    
[📘](https://takeuforward.org/dynamic-programming/matrix-chain-multiplication-dp-48/) [💻](https://www.geeksforgeeks.org/problems/matrix-chain-multiplication0303/1)<br>
