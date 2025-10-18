package controller 

type AuthController struct {}
type AuthControllerInterface interface {}

func NewAuthController() AuthControllerInterface {
	return &AuthController{}
}

func (ac *AuthController) GoogleLogin() {
	
}