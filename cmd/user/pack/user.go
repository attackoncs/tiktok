// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package pack

import (
	"github.com/a76yyyy/tiktok/cmd/user/kitex_gen/user"

	"github.com/a76yyyy/tiktok/cmd/user/dal/db"
)

// User pack user info
func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}

	follow_count := int64(u.FollowerCount)
	follower_count := int64(u.FollowerCount)

	return &user.User{Id: int64(u.ID), Name: u.UserName, FollowCount: &follow_count, FollowerCount: &follower_count}
}

// Users pack list of user info
func Users(us []*db.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}