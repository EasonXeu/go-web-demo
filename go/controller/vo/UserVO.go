package vo

import "go-web-demo/go/dal/dataobject"

type UserVO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" binding:"required"`
	NickName  string `json:"nickName" binding:"required"`
	Avatar    string `json:"avatar"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Mobile    string `json:"mobile"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli"`
}

func UserVO2DO(vo *UserVO) *dataobject.User {
	return &dataobject.User{
		ID:        vo.ID,
		Name:      vo.Name,
		NickName:  vo.NickName,
		Avatar:    vo.Avatar,
		Password:  vo.Password,
		Email:     vo.Email,
		Mobile:    vo.Mobile,
		CreatedAt: vo.CreatedAt,
		UpdatedAt: vo.UpdatedAt,
	}
}

func UserDO2VO(do *dataobject.User) *UserVO {
	return &UserVO{
		ID:        do.ID,
		Name:      do.Name,
		NickName:  do.NickName,
		Avatar:    do.Avatar,
		Password:  do.Password,
		Email:     do.Email,
		Mobile:    do.Mobile,
		CreatedAt: do.CreatedAt,
		UpdatedAt: do.UpdatedAt,
	}
}

func UserDOList2VO(doList []*dataobject.User) []*UserVO {
	var userVOList []*UserVO
	for _, do := range doList {
		userVOList = append(userVOList, UserDO2VO(do))
	}
	return userVOList
}
