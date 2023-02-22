/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-02-22 17:52:32  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-02-22 17:52:32  */
// Runtime22 ms
// Beats
// 95.80%
// Memory7.6 MB
// Beats
// 96.50%

package main

import (
	"container/heap"
)

type KthLargest struct {
	k    int
	heap intHeap
 }
 
 
 func Constructor(k int, nums []int) KthLargest {
	h := intHeap(nums)
	//Convert to heap
	heap.Init(&h)
 
	return KthLargest{
	   k:    k,
	   heap: h,
	}
 }
 
 func (kl *KthLargest) Add(val int) int {
	heap.Push(&kl.heap, val)
 
	//Pop minimum element until len(h) = k and kthlargest = h[0]
	for len(kl.heap) > kl.k {
	   heap.Pop(&kl.heap)
	}
 
	return kl.heap[0]
 }
 
 type intHeap []int
 
 func (h intHeap) Len() int {
	return len(h)
 }
 
 func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
 }
 
 func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
 }
 func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
 }
 
 func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
 }