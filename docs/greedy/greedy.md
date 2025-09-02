## N Meetings in one room

**❓**:  There is one meeting room in a firm. You are given two arrays, start and end each of size N.For an index ‘i’, start[i] denotes the starting time of the ith meeting while end[i]  will denote the ending time of the ith meeting. Find the maximum number of meetings that can be accommodated if only one meeting can happen in the room at a  particular time. Print the order in which these meetings will be performed.<br>
**Example**:  
Input:  N = 6  start = {1,3,0,5,8,5}  end = {2,4,5,7,9,9}  
Output: 1 2 4 5<br>
**🧠**:<br>
1. Sort the items according to end time  
2. Keep track of last meeting end time<br>
2. Iterate through items and check whether the meeting start time is greater than last meeting end time<br>
=== "🐍"
    ```py
    --8<-- "08_greedy/01_nmeetings_in_one_room.py"
    ```
=== "🐍"
    ```py
    --8<-- "08_greedy/01_nmeetings_in_one_room.py"
    ```

[📘](https://takeuforward.org/data-structure/n-meetings-in-one-room/) [💻](https://www.geeksforgeeks.org/problems/n-meetings-in-one-room-1587115620/1)<br>


---

## Minimum number of platforms required for a railway

**❓**: We are given two arrays that represent the arrival and departure times of trains that stop at the platform. We need to find the minimum number of platforms needed at the railway station so that no train has to wait.  
**Example**:  
Input:  N = 6  
arr = {9:00, 9:45, 9:55, 11:00, 15:00, 18:00}  
dep = {9:20, 12:00, 11:30, 11:50, 19:00, 20:00}  
Output: 3  
**🧠**:  
1. Combine both arrival and departure and sort them.  
2. Iterate through both of them, if any arrival increase the counter.  
3. If any departure, decrease the counter.  

=== "🐍"
    ```py
    --8<-- "08_greedy/01_nmeetings_in_one_room.py"
    ```
=== "🐋"
    ```go
    --8<-- "08_greedy/01_nmeetings_in_one_room.go"
    ```

[📘](https://takeuforward.org/data-structure/minimum-number-of-platforms-required-for-a-railway/) [💻](https://www.geeksforgeeks.org/problems/minimum-platforms-1587115620/1)<br>

---

## Job Sequencing Problem

**❓**: You are given a set of N jobs where each job comes with a deadline and profit. The profit can only be earned upon completing the job within its deadline. Find the number of jobs done and the maximum profit that can be obtained. Each job takes a single unit of time and only one job can be performed at a time.  
**Example**:  
Input:  N = 4, Jobs = {(1,4,20),(2,1,10),(3,1,40),(4,1,30)}   
Output: 2 60  
Explanation: The 3rd job with a deadline 1 is performed during the first unit of time .The 1st job is performed during the second unit of time as its deadline is 4.  
Profit = 40 + 20 = 60  
**🧠**:  
1. Sort the jobs pased on profit.  
2. Try to complete the job as late as possible.

=== "🐍"
    ```py
    --8<-- "08_greedy/03_job_sequencing.py"
    ```
=== "🐋"
    ```go
    --8<-- "08_greedy/03_job_sequencing.go"
    ```

[📘](https://takeuforward.org/data-structure/job-sequencing-problem/) [💻](https://www.geeksforgeeks.org/problems/job-sequencing-problem-1587115620/1)<br>

---

## Fractional Knapsack Problem : Greedy Approach

**❓**: The weight of N items and their corresponding values are given. We have to put these items in a knapsack of weight W such that the total value obtained is maximized.  
Note: We can either take the item as a whole or break it into smaller units.    
**Example**:  
Input:  N = 3, W = 50, values[] = {100,60,120}, weight[] = {20,10,30}.  
Output: 240.00  
Explanation: The first and second items  are taken as a whole  while only 20 units of the third item is taken. Total value = 100 + 60 + 80 = 240.00  
**🧠**:  
1. Sort the items according to value(Profit per unit weight).  
2. At the end, when item weight is greater than remaining weight, take the fraction of it and return.  

=== "🐍"
    ```py
    --8<-- "08_greedy/04_fractional_knapsack.py"
    ```
=== "🐋"
    ```go
    --8<-- "08_greedy/04_fractional_knapsack.go"
    ```

[📘](https://takeuforward.org/data-structure/fractional-knapsack-problem-greedy-approach/) [💻](https://www.geeksforgeeks.org/problems/fractional-knapsack-1587115620/1)<br>

---

## Find minimum number of coins

**❓**: Given a value V, if we want to make a change for V Rs, and we have an infinite supply of each of the denominations in Indian currency, i.e., we have an infinite supply of { 1, 2, 5, 10, 20, 50, 100, 500, 1000} valued coins/notes, what is the minimum number of coins and/or notes needed to make the change.     
**Example**:  
Input: V = 70  
Output: 2  
Explaination: We need a 50 Rs note and a 20 Rs note.  
**🧠**:  
Iterate through high value coins to low value coins and try to get as many high value coins as we can.    

=== "🐍"
    ```py
    --8<-- "08_greedy/05_min_coins.py"
    ```
=== "🐋"
    ```go
    --8<-- "08_greedy/05_min_coins.go"
    ```

[📘](https://takeuforward.org/data-structure/find-minimum-number-of-coins/)<br>

---

## Assign Cookies

**❓**: Assume you are an awesome parent and want to give your children some cookies. But, you should give each child at most one cookie.  

Each child i has a greed factor g[i], which is the minimum size of a cookie that the child will be content with; and each cookie j has a size s[j]. If s[j] >= g[i], we can assign the cookie j to the child i, and the child i will be content. Your goal is to maximize the number of your content children and output the maximum number.     
**Example**:  
Input: g = [1,2,3], s = [1,1]  
Output: 1  
Explanation: You have 3 children and 2 cookies. The greed factors of 3 children are 1, 2, 3.   
And even though you have 2 cookies, since their size is both 1, you could only make the child whose greed factor is 1 content.
You need to output 1.   
**🧠**:  
Sort the items and try to satisfy the child with least size cookie possible.
=== "🐍"
    ```py
    --8<-- "08_greedy/06_assign_cookies.py"
    ```
=== "🐋"
    ```go
    --8<-- "08_greedy/06_assign_cookies.go"
    ```

[💻](https://leetcode.com/problems/assign-cookies/)<br>