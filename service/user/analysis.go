package user

type AnalysisService struct{}

type UserAnalysis struct {
}

var ageGroup = map[string]int{
	"<20":   0,
	"20-30": 0,
	"30-40": 0,
	"40-50": 0,
	">50":   0,
}

func (a *AnalysisService) AnalysisUserByAge() {

}
