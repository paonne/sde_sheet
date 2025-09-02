## Next Smaller Element

**🧠**:  
1. Use stack. Iterate through items and remove items from stack whose value is greater than or equal to current num.   
2. Next smaller element is element present in top of stack. If stack is empty, there is not next smaller element.  
=== "🐍"
    ```py
    --8<-- "14_stacks_and_queues_2/01_nse.py"
    ```
=== "🦀"
    ```rust
    --8<-- "14_stacks_and_queues_2/01_nse.rs"
    ```
=== "🐋"
    ```go
    --8<-- "14_stacks_and_queues_2/01_nse.go"
    ```

[💻](https://www.interviewbit.com/problems/nearest-smaller-element/)<br>

---

## Largest rectangle in a histogram

**❓**: Given an array of integers heights representing the histogram's bar height where the width of each bar is 1  return the area of the largest rectangle in histogram.  
**Example**:  
Input: heights = [2,1,5,6,2,3]
Output: 10
Explanation: The above is a histogram where width of each bar is 1.
The largest rectangle is shown in the red area, which has an area = 10 units.

![img](https://assets.leetcode.com/uploads/2021/01/04/histogram.jpg)

**🧠**:  
For every height, possible rectangle width is `nse-pse-1`
=== "🐍"
    ```py
    --8<-- "14_stacks_and_queues_2/04_largest_rectangle_in_histogram.py"
    ```
=== "🦀"
    ```rust
    --8<-- "14_stacks_and_queues_2/04_largest_rectangle_in_histogram.rs"
    ```
=== "🐋"
    ```go
    --8<-- "14_stacks_and_queues_2/04_largest_rectangle_in_histogram.go"
    ```

[📘](https://takeuforward.org/data-structure/area-of-largest-rectangle-in-histogram/) [💻](https://leetcode.com/problems/largest-rectangle-in-histogram/)<br>

---

## Sliding Window Maximum

**❓**: Given an array of integers arr, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

**Example**:  
Input: arr = [4,0,-1,3,5,3,6,8], k = 3

Output: [4,3,5,5,6,8]

Explanation: 

Window position                   Max  
------------------------         -----  
[4  0  -1] 3  5  3  6  8           - 4  
 4 [0  -1  3] 5  3  6  8           - 3  
 4  0 [-1  3  5] 3  6  8           - 5  
 4  0  -1 [3  5  3] 6  8           - 5  
 4  0  -1  3 [5  3  6] 8           - 6  
 4  0  -1  3  5 [3  6  8]          - 8  

For each window of size k=3, we find the maximum element in the window and add it to our output array.
=== "🐍"
    ```py
    --8<-- "14_stacks_and_queues_2/05_sliding_window_maximum.py"
    ```
=== "🦀"
    ```rust
    --8<-- "14_stacks_and_queues_2/05_sliding_window_maximum.rs"
    ```
=== "🐋"
    ```go
    --8<-- "14_stacks_and_queues_2/05_sliding_window_maximum.go"
    ```

[📘](https://takeuforward.org/data-structure/sliding-window-maximum/) [💻](https://leetcode.com/problems/sliding-window-maximum/description/)<br>

---

## Implement Min Stack

**❓**: Implement Min Stack | O(2N) and O(N) Space Complexity. Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

**Example**:  
Input Format:["MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"]
[
[ ], [-2], [0], [-3], [ ], [ ], [ ], [ ]
]

Result: [null, null, null, null, -3, null, 0, -2]
=== "🐍"
    ```py
    --8<-- "14_stacks_and_queues_2/06_min_stack.py"
    ```
=== "🦀"
    ```rust
    --8<-- "14_stacks_and_queues_2/06_min_stack.rs"
    ```
=== "🐋"
    ```go
    --8<-- "14_stacks_and_queues_2/06_min_stack.go"
    ```

[📘](https://takeuforward.org/data-structure/implement-min-stack-o2n-and-on-space-complexity/) [💻](https://leetcode.com/problems/min-stack/description/)<br>

---

## Rotten Orange (Using BFS)

**❓**: You will be given an m x n grid, where each cell has the following values : 

2  -  represents a rotten orange
1  -  represents a Fresh orange
0  -  represents an Empty Cell
Every minute, if a Fresh Orange is adjacent to a Rotten Orange in 4-direction ( upward, downwards, right, and left ) it becomes Rotten. 

Return the minimum number of minutes required such that none of the cells has a Fresh Orange. If it's not possible, return -1.

**Example**:  

Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
Output: 4

![img](https://assets.leetcode.com/uploads/2019/02/16/oranges.png)

**🧠**:  
Use BFS
=== "🐍"
    ```py
    --8<-- "14_stacks_and_queues_2/07_rotten_oranges.py"
    ```
=== "🦀"
    ```rust
    --8<-- "14_stacks_and_queues_2/07_rotten_oranges.rs"
    ```
=== "🐋"
    ```go
    --8<-- "14_stacks_and_queues_2/07_rotten_oranges.go"
    ```

[📘](https://takeuforward.org/data-structure/rotten-oranges-min-time-to-rot-all-oranges-bfs/) [💻](https://leetcode.com/problems/rotting-oranges/description/)<br>

---

## Stock span problem

**❓**: Design an algorithm that collects daily price quotes for some stock and returns the span of that stock's price for the current day.

The span of the stock's price in one day is the maximum number of consecutive days (starting from that day and going backward) for which the stock price was less than or equal to the price of that day.

For example, if the prices of the stock in the last four days is [7,2,1,2] and the price of the stock today is 2, then the span of today is 4 because starting from today, the price of the stock was less than or equal 2 for 4 consecutive days.

Implement the StockSpanner class:

StockSpanner() Initializes the object of the class.
int next(int price) Returns the span of the stock's price given that today's price is price.

**Example**:  
Input
["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
[[], [100], [80], [60], [70], [60], [75], [85]]
Output
[null, 1, 1, 1, 2, 1, 4, 6]

**🧠**:  
Find pge
=== "🐍"
    ```py
    --8<-- "14_stacks_and_queues_2/08_stock_span.py"
    ```
=== "🦀"
    ```rust
    --8<-- "14_stacks_and_queues_2/08_stock_span.rs"
    ```
=== "🐋"
    ```go
    --8<-- "14_stacks_and_queues_2/08_stock_span.go"
    ```

[💻](https://leetcode.com/problems/online-stock-span/description/)<br>

---

## The Celebrity Problem

**❓**: A celebrity is a person who is known by everyone else at the party but does not know anyone in return. Given a square matrix M of size N x N where M[i][j] is 1 if person i knows person j, and 0 otherwise, determine if there is a celebrity at the party. Return the index of the celebrity or -1 if no such person exists.

Note that M[i][i] is always 0.

**Example**:  
Input: M = [ [0, 1, 1, 0], [0, 0, 0, 0], [1, 1, 0, 0], [0, 1, 1, 0] ]

Output: 1

Explanation: Person 1 does not know anyone and is known by persons 0, 2, and 3. Therefore, person 1 is the celebrity.
=== "🐍"
    ```py
    --8<-- "14_stacks_and_queues_2/10_celebrity.py"
    ```
=== "🦀"
    ```rust
    --8<-- "14_stacks_and_queues_2/10_celebrity.rs"
    ```
=== "🐋"
    ```go
    --8<-- "14_stacks_and_queues_2/10_celebrity.go"
    ```

---