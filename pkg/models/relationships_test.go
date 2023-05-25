package models

import "testing"

func TestRelationship_CreateRelationship(t *testing.T) {
	/*relation := []Relationship{
		{
			ID:             5,
			FollowedUserId: 1,
			FollowerUserId: 2,
		},
		{
			ID:             6,
			FollowedUserId: 1,
			FollowerUserId: 3,
		},
		{
			ID:             7,
			FollowedUserId: 3,
			FollowerUserId: 2,
		},
		{
			ID:             8,
			FollowedUserId: 2,
			FollowerUserId: 1,
		},
		{
			ID:             9,
			FollowedUserId: 4,
			FollowerUserId: 2,
		},
	}*/
	//for _, v := range relation {
	//	err := v.CreateRelationship()
	//	if err != nil {
	//		t.Error("can't create relationship,err:", err)
	//		return
	//	} else {
	//		t.Log("创建关系成功！")
	//	}
	//}

	r2, err := GetFollowed(2)
	r3, err := GetFollowers(2)
	t.Log(err)
	t.Logf("2关注的所有人有:%v\n 关注2的所有人是:%v", r2, r3)
}
