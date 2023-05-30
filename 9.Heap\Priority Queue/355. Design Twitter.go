package main

import (
	"container/heap"
)

type MyHeap [][4]int

func (h MyHeap) Len() int           { return len(h) }
func (h MyHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h MyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MyHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([4]int))
}

func (h *MyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MyHeap) Top() interface{} {
	return h[0]
}

type Twitter struct {
	cnt int
	// userId -> to list tweets
	tweetMap map[int][][2]int
	// userId -> to list follewees
	followedMap map[int]map[int]bool
}

func Constructor() Twitter {
	return Twitter{
		cnt:         0,
		tweetMap:    map[int][][2]int{},
		followedMap: map[int]map[int]bool{},
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweetMap[userId] = append(this.tweetMap[userId], [2]int{this.cnt, tweetId})
	this.cnt++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	res := []int{}
	maxHeap := &MyHeap{}
	heap.Init(maxHeap)

    // add following them selve
	if this.followedMap[userId] == nil {
		this.followedMap[userId] = map[int]bool{}
	}
    this.followedMap[userId][userId] = true
    
	followees := this.followedMap[userId]
	for followeeId, _ := range followees {
		if listTweets, found := this.tweetMap[followeeId]; found {
			// listTweets := this.tweetMap[followeeId]
			index := len(listTweets) - 1
			count, tweetID := listTweets[index][0], listTweets[index][1]
			// push tweetID to max heap
			heap.Push(maxHeap, [4]int{count, tweetID, followeeId, index - 1})
		}
	}

	for maxHeap.Len() > 0 && len(res) < 10 {
		tweet := heap.Pop(maxHeap).([4]int)
		_, tweetID, followeeId, index := tweet[0], tweet[1], tweet[2], tweet[3]
		res = append(res, tweetID)
		if index >= 0 {
			prevTweet := this.tweetMap[followeeId][index]
			prevCount, prevTweetID := prevTweet[0], prevTweet[1]
			heap.Push(maxHeap, [4]int{prevCount, prevTweetID, followeeId, index - 1})
		}
	}

	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if v, found := this.followedMap[followerId]; found {
		v[followeeId] = true
	} else {
		this.followedMap[followerId] = map[int]bool{}
		this.followedMap[followerId][followeeId] = true
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if v, found := this.followedMap[followerId]; found {
		delete(v, followeeId)
	}
}
