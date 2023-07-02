package modules

type Semeru3Controller struct {
	uc Semeru3Usecase
}

func CreateSemeru3Controller(uc Semeru3Usecase) Semeru3Controller {
	return Semeru3Controller{
		uc: uc,
	}
}
