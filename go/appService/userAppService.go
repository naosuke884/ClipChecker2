package appservice

//user登録

//user認証→share/auth.go
//帰り値はUser構造体

//providerがtwitchの場合→Infrastructure/provider/twitchAuth.go
//twitchAuthで、Userを認証（tokenの獲得）
//sessionで認証状態を管理('isAuth')
//tokenからTwitchサーバーからUser情報を獲得（Name,Email,ImageURL）
//Userで返す

//Userを永続化→Domain/user.go

//driverがsqlite3の場合→Infrastructure/driver/sqlite3.go

type UserAppService struct {
}

func NewUserAppService() *UserAppService {
	return &UserAppService{}
}

func (u *UserAppService) Register() {
}

func (u *UserAppService) Login() {
}

func (u *UserAppService) Logout() {
}

func (u *UserAppService) Get() {
}

func (u *UserAppService) Update() {
}

func (u *UserAppService) Delete() {
}
