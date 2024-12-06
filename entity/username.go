package entity


type UserNameList []string

type IUserName interface {
	IsValidUserName(username string) bool
}

func (u UserNameList) IsValidUserName(username string) bool {
	for _, userNm := range u {
		if userNm == username {
			return false
		}
	}
	return true
}