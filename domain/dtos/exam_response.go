package dtos

type ExamsResponse struct {
	Exams []*ExamResponse `json:"exams"`
}

type ExamResponse struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}
