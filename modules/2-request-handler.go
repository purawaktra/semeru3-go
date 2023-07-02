package modules

type Semeru3RequestHandler struct {
	ctrl Semeru3Controller
}

func CreateSemeru3RequestHandler(ctrl Semeru3Controller) Semeru3RequestHandler {
	return Semeru3RequestHandler{
		ctrl: ctrl,
	}
}
