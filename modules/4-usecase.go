package modules

type Semeru3Usecase struct {
	repo Semeru3Repo
}

func CreateSemeru3Usecase(repo Semeru3Repo) Semeru3Usecase {
	return Semeru3Usecase{
		repo: repo,
	}
}
