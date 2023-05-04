package section

import (
	"fmt"
	"time"
)

/*
 * @lc app=leetcode.cn id=355 lang=golang
 *
 * [355] 设计推特
 */

// @lc code=start

type TwitterItem struct {
	Id        int
	CreatedAt time.Time
}

type Twitter struct {
	UserMaps   map[int][]TwitterItem
	FollowMaps map[int]map[int]struct{}
}

func Constructor355() Twitter {
	return Twitter{
		UserMaps:   make(map[int][]TwitterItem, 0),
		FollowMaps: make(map[int]map[int]struct{}, 0),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.UserMaps[userId] = append(this.UserMaps[userId], TwitterItem{
		Id:        tweetId,
		CreatedAt: time.Now(),
	})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	twitterLists := [][]TwitterItem{this.UserMaps[userId]}
	for followeeId := range this.FollowMaps[userId] {
		twitterLists = append(twitterLists, this.UserMaps[followeeId])
	}

	results := make([]int, 0)
	for i := 0; i < 10; i++ {
		chooseIndex := -1
		chooseTime := time.Time{}
		for j := 0; j < len(twitterLists); j++ {
			if len(twitterLists[j]) > 0 {
				if twitterLists[j][len(twitterLists[j])-1].CreatedAt.After(chooseTime) {
					chooseIndex = j
					chooseTime = twitterLists[j][len(twitterLists[j])-1].CreatedAt
				}
			}
		}

		if chooseTime.IsZero() {
			break
		}

		results = append(results, twitterLists[chooseIndex][len(twitterLists[chooseIndex])-1].Id)
		twitterLists[chooseIndex] = twitterLists[chooseIndex][:len(twitterLists[chooseIndex])-1]
	}

	return results
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.FollowMaps[followerId]; !ok {
		this.FollowMaps[followerId] = make(map[int]struct{}, 0)
	}

	this.FollowMaps[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.FollowMaps[followerId]; !ok {
		this.FollowMaps[followerId] = make(map[int]struct{}, 0)
	}

	delete(this.FollowMaps[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// @lc code=end

func Test355() {
	twitter := Constructor355()
	twitter.PostTweet(1, 5) // 用户 1 发送了一条新推文 (用户 id = 1, 推文 id = 5)
	// [5]
	fmt.Println(twitter.GetNewsFeed(1)) // 用户 1 的获取推文应当返回一个列表，其中包含一个 id 为 5 的推文
	twitter.Follow(1, 2)                // 用户 1 关注了用户 2
	twitter.PostTweet(2, 6)             // 用户 2 发送了一个新推文 (推文 id = 6)
	// [6, 5]
	fmt.Println(twitter.GetNewsFeed(1)) // 用户 1 的获取推文应当返回一个列表，其中包含两个推文，id 分别为 -> [6, 5] 。推文 id 6 应当在推文 id 5 之前，因为它是在 5 之后发送的
	twitter.Unfollow(1, 2)              // 用户 1 取消关注了用户 2
	// [5]
	fmt.Println(twitter.GetNewsFeed(1)) // 用户 1 获取推文应当返回一个列表，其中包含一个 id 为 5 的推文。因为用户 1 已经不再关注用户 2
}
