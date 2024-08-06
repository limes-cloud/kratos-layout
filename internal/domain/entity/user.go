package entity

type User struct {
	Id        uint32  `json:"id,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Email     *string `json:"email,omitempty"`
	Username  *string `json:"username,omitempty"`
	NickName  string  `json:"nickName,omitempty"`
	RealName  *string `json:"realName,omitempty"`
	Avatar    *string `json:"avatar,omitempty"`
	AvatarUrl *string `json:"avatarUrl,omitempty"`
}
