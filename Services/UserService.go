package Services

type IUserService interface {
	GetName(userId int) string
}

//实现类
type UserService struct {

}

func (this UserService) GetName(userId int) string {
	if userId == 101{
		return "LinCG"
	}
	return "guest"
}