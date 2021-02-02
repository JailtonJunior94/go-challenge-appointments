package facades

import (
	"encoding/json"
	"net/http"

	"github.com/jailtonjunior94/go-challenge-appointments/domain/dtos"
	"github.com/jailtonjunior94/go-challenge-appointments/infrastructure/environments"
)

type IExamFacade interface {
	GetExams() (exams *dtos.ExamsResponse, err error)
}

type ExamFacade struct {
}

func NewExamFacade() IExamFacade {
	return &ExamFacade{}
}

func (e ExamFacade) GetExams() (exams *dtos.ExamsResponse, err error) {
	response, err := http.Get(environments.ExamsBaseUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&exams); err != nil {
		return nil, err
	}

	return exams, nil
}
