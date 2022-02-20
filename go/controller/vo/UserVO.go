package vo

type UserVO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	NickName  string `json:"nickName"`
	Avatar    string `json:"avatar"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli"`
}
