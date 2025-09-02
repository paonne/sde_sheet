## Level order Traversal

**Example**:  

![alternate](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)

Input: root = [3,9,20,null,null,15,7]  
Output: [[3],[9,20],[15,7]]  

**🧠**:  
BFS
=== "🐍"
    ```py
    --8<-- "18_binary_tree_ii/01_level_order_traversal.py"
    ```
=== "🐋"
    ```go
    --8<-- "18_binary_tree_ii/01_level_order_traversal.go"
    ```

[📘](https://takeuforward.org/data-structure/level-order-traversal-of-a-binary-tree/) [💻](https://leetcode.com/problems/binary-tree-level-order-traversal/description/)<br>

---

## Height of Binary Tree
 
**Example**:  

![alternate](https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg)  
Input: root = [3,9,20,null,null,15,7]  
Output: 3
=== "🐍"
    ```py
    --8<-- "18_binary_tree_ii/02_height_of_bt.py"
    ```
=== "🐋"
    ```go
    --8<-- "18_binary_tree_ii/02_height_of_bt.go"
    ```

[📘](https://takeuforward.org/data-structure/maximum-depth-of-a-binary-tree/) [💻](https://leetcode.com/problems/maximum-depth-of-binary-tree/description/)<br>

---

## Diamter of Binary Tree

**Example**:  

![alternate](https://assets.leetcode.com/uploads/2021/03/06/diamtree.jpg)  
Input: root = [1,2,3,4,5]  
Output: 3  
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].

=== "🐍"
    ```py
    --8<-- "18_binary_tree_ii/03_diameter_of_bt.py"
    ```
=== "🐋"
    ```go
    --8<-- "18_binary_tree_ii/03_diameter_of_bt.go"
    ```

[📘](https://takeuforward.org/data-structure/calculate-the-diameter-of-a-binary-tree/) [💻](https://leetcode.com/problems/diameter-of-binary-tree/)<br>

---

## Balanced Binary Tree

**❓**: Given a binary tree, determine if it is height-balanced

A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

=== "🐍"
    ```py
    --8<-- "18_binary_tree_ii/04_balanced_bt.py"
    ```
=== "🐋"
    ```go
    --8<-- "18_binary_tree_ii/04_balanced_bt.go"
    ```

[📘](https://takeuforward.org/data-structure/check-if-the-binary-tree-is-balanced-binary-tree/) [💻](https://leetcode.com/problems/balanced-binary-tree/)<br>

---

## LCA
 
**Example**:  

![alternate](https://assets.leetcode.com/uploads/2018/12/14/binarytree.png)   
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1  
Output: 3  
Explanation: The LCA of nodes 5 and 1 is 3.

**🧠**:  
1. Return None or Node in every recusive call for both left and right childs.  
2. If both left and right recursive calls is not None, then current node is your answer.  

=== "🐍"
    ```py
    --8<-- "18_binary_tree_ii/05_lca.py"
    ```
=== "🐋"
    ```go
    --8<-- "18_binary_tree_ii/05_lca.go"
    ```

[📘](https://takeuforward.org/data-structure/lowest-common-ancestor-for-two-given-nodes/) [💻](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)<br>

---

## Check if two trees are identical or not

=== "🐍"
    ```py
    --8<-- "18_binary_tree_ii/06_identical_trees.py"
    ```
=== "🐋"
    ```go
    --8<-- "18_binary_tree_ii/06_identical_trees.go"
    ```

[📘](https://takeuforward.org/data-structure/check-if-two-trees-are-identical/) [💻](https://leetcode.com/problems/same-tree/)<br>

---

## Zigzag Level Order Traversal

**Example**:  

![alternate](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)  
Input: root = [3,9,20,null,null,15,7]  
Output: [[3],[20,9],[15,7]]  

**🧠**:  
1. Use BFS  
2. Use zig_zag flag at every level to zig zag the iteration.  

=== "🐍"
    ```py
    --8<-- "18_binary_tree_ii/07_zig_zag_traversal.py"
    ```
=== "🐋"
    ```go
    --8<-- "18_binary_tree_ii/07_zig_zag_traversal.go"
    ```

[📘](https://takeuforward.org/data-structure/zig-zag-traversal-of-binary-tree/) [💻](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/)<br>

---

## Boundary Traversal

**❓**: Given a Binary Tree, perform the boundary traversal of the tree. The boundary traversal is the process of visiting the boundary nodes of the binary tree in the anticlockwise direction, starting from the root.

**Example**:  

![alternate](https://static.takeuforward.org/content/boundary-traversal-image1-XAwduImr){ width=30% }  
Input:Binary Tree: 1 2 7 3 -1 -1 8 -1 4 9 -1 5 6 10 11  
Output: Boundary Traversal: [1, 2, 3, 4, 5, 6, 10, 11, 9, 8, 7]  

**🧠**:  
1. Result is - root + left boundary excluding leaves + leaf nodes + reversed right boundary excluding leaves  
2. Edge case - If root is leave node, return [root value]  

```py
{!18_binary_tree_ii/08_boundary_traversal.py!}
```

[📘](https://takeuforward.org/data-structure/boundary-traversal-of-a-binary-tree/) [💻](https://www.naukri.com/code360/problems/boundary-traversal-of-binary-tree_790725?leftPanelTabValue=SUBMISSION)<br>

---