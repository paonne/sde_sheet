## Minimum Path Sum

**❓**: Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.  

**Example**:  
Input: grid = [[1,3,1],[1,5,1],[4,2,1]]  
Output: 7  
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.  

=== "🐍"
    ```py
    --8<-- "26_dp_ii/01_min_path_sum.py"
    ```
=== "🦀"
    ```rust
    --8<-- "26_dp_ii/01_min_path_sum.rs"
    ```
=== "🐋"
    ```go
    --8<-- "26_dp_ii/01_min_path_sum.go"
    ```
    
[📘](https://takeuforward.org/data-structure/minimum-path-sum-in-a-grid-dp-10/) [💻](https://leetcode.com/problems/minimum-path-sum/description/)<br>

## Coin Change

**❓**: You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.  

**Example**:  
Input: coins = [1,2,5], amount = 11  
Output: 3  
Explanation: 11 = 5 + 5 + 1  

=== "🐍"
    ```py
    --8<-- "26_dp_ii/02_coin_change.py"
    ```
=== "🦀"
    ```rust
    --8<-- "26_dp_ii/02_coin_change.rs"
    ```
=== "🐋"
    ```go
    --8<-- "26_dp_ii/02_coin_change.go"
    ```
    
[📘](https://takeuforward.org/data-structure/coin-change-2-dp-22/) [💻](https://leetcode.com/problems/coin-change/description/)<br>

## Partition Equal Subset Sum

**❓**: Given an integer array nums, return true if you can partition the array into two subsets such that the sum of the elements in both subsets is equal or false otherwise.

**Example**:  
Input: nums = [1,5,11,5]  
Output: true  
Explanation: The array can be partitioned as [1, 5, 5] and [11].  

=== "🐍"
    ```py
    --8<-- "26_dp_ii/03_subset_sum.py"
    ```
=== "🦀"
    ```rust
    --8<-- "26_dp_ii/03_subset_sum.rs"
    ```
=== "🐋"
    ```go
    --8<-- "26_dp_ii/03_subset_sum.go"
    ```
    
[📘](https://takeuforward.org/data-structure/subset-sum-equal-to-target-dp-14/) [💻](https://leetcode.com/problems/partition-equal-subset-sum/)<br>

## Rod Cutting

**❓**: Given a wooden stick of length n units. The stick is labelled from 0 to n. For example, a stick of length 6 is labelled as follows:

Given an integer array cuts where cuts[i] denotes a position you should perform a cut at.

You should perform the cuts in order, you can change the order of the cuts as you wish.

The cost of one cut is the length of the stick to be cut, the total cost is the sum of costs of all cuts. When you cut a stick, it will be split into two smaller sticks (i.e. the sum of their lengths is the length of the stick before the cut). Please refer to the first example for a better explanation.

Return the minimum total cost of the cuts.

Problem statement is confusing, please refer leetcode link for better clarity of description.  

=== "🐍"
    ```py
    --8<-- "26_dp_ii/04_rod_cutting.py"
    ```
=== "🦀"
    ```rust
    --8<-- "26_dp_ii/04_rod_cutting.rs"
    ```
=== "🐋"
    ```go
    --8<-- "26_dp_ii/04_rod_cutting.go"
    ```
    
[📘](hhttps://takeuforward.org/data-structure/rod-cutting-problem-dp-24/) [💻](https://leetcode.com/problems/minimum-cost-to-cut-a-stick/description/)<br>

## Egg Dropping

**❓**: You are given n identical eggs and you have access to a k-floored building from 1 to k.

There exists a floor f where 0 <= f <= k such that any egg dropped from a floor higher than f will break, and any egg dropped from or below floor f will not break.
There are few rules given below.  
1. An egg that survives a fall can be used again.  
2. A broken egg must be discarded.  
3. The effect of a fall is the same for all eggs.  
4. If the egg doesn't break at a certain floor, it will not break at any floor below.  
5. If the egg breaks on a certain floor, it will break on any floor above.  

Return the minimum number of moves you need to determine the value of f with certainty.  

**Example**:  
Input: n = 1, k = 2  
Output: 2  
Explanation: Drop the egg from floor 1. If it breaks, we know that f = 0. Otherwise, drop the egg from floor 2.If it breaks, we know that f = 1.  If it does not break, then we know f = 2. Hence, we need at minimum 2 moves to determine with certainty what the value of f is.  

Input: n = 10, k = 5  
Output: 3  
Explanation: Drop the egg from floor 2. If it breaks, test floor 1 with a remaining egg.If it doesn’t break, drop from floor 4. If it breaks, test floor 3. If it still doesn’t break, we know the critical floor is 5.Hence, with a minimum of 3 moves, we can find the critical floor.  

=== "🐍"
    ```py
    --8<-- "26_dp_ii/05_egg_dropping.py"
    ```
=== "🦀"
    ```rust
    --8<-- "26_dp_ii/05_egg_dropping.rs"
    ```
    
[💻](https://www.geeksforgeeks.org/problems/egg-dropping-puzzle-1587115620/1)<br>

## Word Break

**❓**: Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.  

**Example**:  
Input: s = "leetcode", wordDict = ["leet","code"]  
Output: true  
Explanation: Return true because "leetcode" can be segmented as "leet code".  

Input: s = "applepenapple", wordDict = ["apple","pen"]  
Output: true  
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.  

=== "🐍"
    ```py
    --8<-- "26_dp_ii/06_word_break.py"
    ```
=== "🦀"
    ```rust
    --8<-- "26_dp_ii/06_word_break.rs"
    ```
=== "🐋"
    ```go
    --8<-- "26_dp_ii/06_word_break.go"
    ```
    
[💻](https://leetcode.com/problems/word-break/)<br>

## Palindrome Partitioning

**❓**: Given a string s, a partitioning of the string is a palindrome partitioning if every sub-string of the partition is a palindrome. Determine the fewest cuts needed for palindrome partitioning of the given string.  

**Example**:  
Input: s = "ababbbabbababa"  
Output: 3  
Explaination: After 3 partitioning substrings 
are "a", "babbbab", "b", "ababa".  

Input: s = "aaabba"  
Output: 1  
Explaination: The substrings after 1 partitioning are "aa" and "abba".  

=== "🐍"
    ```py
    --8<-- "26_dp_ii/07_palindrome_partitioning.py"
    ```
=== "🦀"
    ```rust
    --8<-- "26_dp_ii/07_palindrome_partitioning.rs"
    ```
=== "🐋"
    ```go
    --8<-- "26_dp_ii/07_palindrome_partitioning.go"
    ```
    
[💻](https://www.geeksforgeeks.org/problems/palindromic-patitioning4845/1)<br>

## Job Sequencing Problem

**❓**: You are given three arrays: id, deadline, and profit, where each job is associated with an ID, a deadline, and a profit. Each job takes 1 unit of time to complete, and only one job can be scheduled at a time. You will earn the profit associated with a job only if it is completed by its deadline.

Your task is to find: The maximum number of jobs that can be completed within their deadlines. The total maximum profit earned by completing those jobs.  

**Example**:  
Input: id = [1, 2, 3, 4], deadline = [4, 1, 1, 1], profit = [20, 1, 40, 30]  
Output: [2, 60]  
Explanation: Job1 and Job3 can be done with maximum profit of 60 (20+40).  

Input: id = [1, 2, 3, 4, 5], deadline = [2, 1, 2, 1, 1], profit = [100, 19, 27, 25, 15]  
Output: [2, 127]  
Explanation: Job1 and Job3 can be done with maximum profit of 127 (100+27).  

=== "🐍"
    ```py
    --8<-- "26_dp_ii/08_job_scheduling.py"
    ```
    
[💻](https://www.geeksforgeeks.org/problems/job-sequencing-problem-1587115620/1)<br>
