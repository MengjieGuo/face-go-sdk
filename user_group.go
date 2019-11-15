package face

import (
	"github.com/comeonjy/face-go-sdk/lib"
	"net/url"
)

// 添加用户组
func (f *Face) AddUserGroup(group_id string) (err error) {
	err = f.tool(lib.USER_GROUP_ADD, url.Values{
		"group_id": {group_id},
	}, &Reply{})
	return
}

// 删除用户组
func (f *Face) DelUserGroup(group_id string) (err error) {
	err = f.tool(lib.USER_GROUP_DEL, url.Values{
		"group_id": {group_id},
	}, &Reply{})
	return
}

type GroupList struct {
	GroupIdList []string `json:"group_id_list"`
}

// 用户组列表
func (f *Face) ListUserGroup() (res *GroupList, err error) {
	res = &GroupList{}
	err = f.tool(lib.USER_GROUP_LIST, url.Values{}, &Reply{
		Result: res,
	})
	return
}
