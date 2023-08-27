/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-04-05 20:30:16  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-04-05 20:30:16  */
// Runtime533 ms
// Beats
// 92.18%
// Memory87.1 MB
// Beats
// 12.96%

package main

func main() {
	
}

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
    Head *Node
	Tail *Node
	HT map[int]*Node
	Cap int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
		HT: make(map[int]*Node),
		Cap: capacity,
	}
}


func (this *LRUCache) Get(key int) int {
    node,ok := this.HT[key]
	if ok {
		this.Remove(node)
		this.Add(node)
		return node.value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	node,ok := this.HT[key]
	if ok {
		node.value = value
		this.Remove(node)
		this.Add(node)
		return 
	} else {
		node = &Node{key: key ,  value:value}
		this.HT[key] = node
		this.Add(node)
	}
	if len(this.HT) > this.Cap {
		delete(this.HT, this.Tail.key)
		this.Remove(this.Tail)
	}
    
}

func (this *LRUCache) Add(node *Node) {
	node.prev = nil
	node.next = this.Head
	if this.Head != nil {
		this.Head.prev = node
	}
	this.Head = node
	if this.Tail == nil {
		this.Tail = node
	}

}
func (this *LRUCache) Remove(node *Node) {
	if node != this.Head {
		node.prev.next = node.next
	} else {
		this.Head = node.next
	}
	if node != this.Tail {
		node.next.prev = node.prev
	} else {
		this.Tail = node.prev
	}
}