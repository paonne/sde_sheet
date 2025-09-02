## Maximum path sum

**❓**: A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

**Example**:

![alternate](https://assets.leetcode.com/uploads/2020/10/13/exx2.jpg)  
Input: root = [-10,9,20,null,null,15,7]  
Output: 42  
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.   

=== "🐍"
    ```py
    --8<-- "19_binary_tree_iii/01_max_path_sum.py"
    ```
=== "🐋"
    ```go
    --8<-- "19_binary_tree_iii/01_max_path_sum.go"
    ```

[📘](https://takeuforward.org/data-structure/maximum-sum-path-in-binary-tree/) [💻](https://leetcode.com/problems/binary-tree-maximum-path-sum/)<br>

---

## Construct Binary Tree from inorder and preorder

=== "🐍"
    ```py
    --8<-- "19_binary_tree_iii/02_tree_from_inorder_preorder.py"
    ```
=== "🐋"
    ```go
    --8<-- "19_binary_tree_iii/02_tree_from_inorder_preorder.go"
    ```

[📘](https://takeuforward.org/data-structure/construct-a-binary-tree-from-inorder-and-preorder-traversal/) [💻](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)<br>

---

## Construct Binary Tree from Inorder and Postorder

=== "🐍"
    ```py
    --8<-- "19_binary_tree_iii/03_tree_from_inorder_postorder.py"
    ```
=== "🐋"
    ```go
    --8<-- "19_binary_tree_iii/03_tree_from_inorder_postorder.go"
    ```

[💻](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)<br>

---

## Symmetric Binary Tree

=== "🐍"
    ```py
    --8<-- "19_binary_tree_iii/04_symmetric_bt.py"
    ```
=== "🐋"
    ```go
    --8<-- "19_binary_tree_iii/04_symmetric_bt.go"
    ```

[📘](https://takeuforward.org/data-structure/check-for-symmetrical-binary-tree/) [💻](https://leetcode.com/problems/symmetric-tree/)<br>

---

## Mirror Tree

Given a Binary Tree, convert it into its mirror.  

![alternate](https://contribute.geeksforgeeks.org/wp-content/uploads/mirrortrees.jpg){ width=50%}

```py
{!19_binary_tree_iii/06_is_mirror.py!}
```

[💻](https://practice.geeksforgeeks.org/problems/mirror-tree/1)<br>

---

## Check for Children Sum Property

**❓**: Given a binary tree having n nodes. Check whether all of its nodes have a value equal to the sum of their child nodes. Return 1 if all the nodes in the tree satisfy the given properties, else it returns 0. For every node, the data value must be equal to the sum of the data values in the left and right children. Consider the data value 0 for a NULL child. Also, leaves are considered to follow the property.

**Example**:  
```
       35
      /  \
     20   15
    / \   / \
   15  5 10  5
```
Output: 1  
Explanation: Here, every node is sum of its left and right child.  

```py
{!19_binary_tree_iii/07_check_for_children_sum.py!}
```

[📘](https://takeuforward.org/data-structure/check-for-children-sum-property-in-a-binary-tree/) [💻](https://www.geeksforgeeks.org/problems/children-sum-parent/1)<br>

---