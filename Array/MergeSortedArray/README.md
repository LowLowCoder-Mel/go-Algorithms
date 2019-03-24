# Merge Sorted Array（合并两个有序数组）

**LeetCode 88**

- [英文版](https://leetcode.com/problems/merge-sorted-array/)
- [中文版](https://leetcode-cn.com/problems/merge-sorted-array/)

## 题目
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:

- 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
- 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

示例:
```
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
```

## 思路
既然两个数组都是有序数组,那么就用一个循环比较将小值放入临时数组内,当较短数组结束后,再将较长数组追加到临时数组后面,再将整个数组复制到原来到数组里面