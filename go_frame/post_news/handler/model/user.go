package model

type LoginUserReq struct {
	Name     string `json:"name" binding:"required,min=1,max=10"`
	Password string `json:"password" binding:"required,len=32"`
}

type ModifyPassReq struct {
	// Uid         int    `json:"uid" binding:"required"`
	OldPassword string `json:"old_password" binding:"required,len=32"`
	NewPassword string `json:"new_password" binding:"required,len=32"`
}
