package conf


type Env struct {
	ServerPort   string
	FaceIOAppCode string
	CodeErr int
	CodeSuccess int
	MsgErr string
	MsgSuccess string
	LoginName string
	LoginPassWord string
	LoginUserFacialId string
	CryptoKey string
}

func GetEnv() *Env {
	return &env
}