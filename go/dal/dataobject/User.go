package dataobject

func (User) TableName() string {
	return "sys_user"
}

type User struct {
	ID        uint   `gorm:"column:id; primaryKey; autoIncrement"`
	Name      string `gorm:"column:name; type:varchar(50); not null; uniqueIndex"`
	NickName  string `gorm:"column:nick_name; type:varchar(50);"`
	Avatar    string `gorm:"column:avatar; type:varchar(200);"`
	Password  string `gorm:"column:password; type:varchar(20);"`
	Email     string `gorm:"column:email; type:varchar(100);"`
	Mobile    string `gorm:"column:mobile; type:varchar(20);"`
	CreatedAt int64  `gorm:"column:created_at; type:bigint; not null; autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at; type:bigint; not null; autoUpdateTime:milli"`
	DeletedAt int64  `gorm:"column:deleted_at; type:bigint;"`
}
