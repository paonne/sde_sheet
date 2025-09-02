## Populate Next Right pointers of Tree

**❓**: You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}

Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL. 

**Example**:  

![alternate](https://assets.leetcode.com/uploads/2019/02/14/116_sample.png)  
Input: root = [1,2,3,4,5,6,7]  
Output: [1,#,2,3,#,4,5,6,7,#]  

**🧠**:  
At each level, for every node, perform 2 steps
    a. make left next to right
    b. If node has any next, make right next as next left

```py
{!20_binary_search_tree/01_populate_next_pointers.py!}
```

[💻](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/)<br>

---

## Construct BST from given keys

=== "🐍"
    ```py
    --8<-- "20_binary_search_tree/03_convert_srted_array_to_bst.py"
    ```
=== "🐋"
    ```go
    --8<-- "20_binary_search_tree/03_convert_srted_array_to_bst.go"
    ```

[💻](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)<br>

---

## Construct a BST from a preorder traversal

**🧠**:  
1. For BST, inorder will be sorted.  
2. Find inorder by sorting preorder and construct BT.  

=== "🐍"
    ```py
    --8<-- "20_binary_search_tree/04_bst_from_preorder.py"
    ```
=== "🐋"
    ```go
    --8<-- "20_binary_search_tree/04_bst_from_preorder.go"
    ```

[💻](https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/)<br>

---

## Check is a BT is BST or not

=== "🐍"
    ```py
    --8<-- "20_binary_search_tree/05_validate_bst.py"
    ```
=== "🐋"
    ```go
    --8<-- "20_binary_search_tree/05_validate_bst.go"
    ```

[💻](https://leetcode.com/problems/validate-binary-search-tree/)<br>

---

## Find LCA of two nodes in BST

=== "🐍"
    ```py
    --8<-- "20_binary_search_tree/06_lca_in_bst.py"
    ```
=== "🐋"
    ```go
    --8<-- "20_binary_search_tree/06_lca_in_bst.go"
    ```

[💻](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)<br>

---

## Find the inorder predecessor/successor of a given Key in BST

**❓**: You are given root node of the BST and the key node of the tree. You need to find the in-order successor and predecessor of a given key. If either predecessor or successor is not found, then set it to NULL.

Note:- In an inorder traversal the number just smaller than the target is the predecessor and the number just greater than the target is the successor. 

**Example**:  
```
      8
    /   \
   1     9
   \      \
    4      10
   /
  3
```
key = 8  
Output: 4 9  

```py
{!20_binary_search_tree/07_inorder_successor_predecessor.py!}
```

[💻](https://practice.geeksforgeeks.org/problems/predecessor-and-successor/1)<br>

---