package controller

type Controller struct {
	UserController *UserController
}

func NewController(userController *UserController) *Controller {
	return &Controller{
		UserController: userController,
	}
}
