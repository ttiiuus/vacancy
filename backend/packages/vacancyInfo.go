package vacancy

import (
	"fmt"
)

type VacancyStruct struct {
	ID            string   `json:"id"`
	FirstName     string   `json:"firstName"`
	SecondName    string   `json:"secondName"`
	JobExperience int      `json:"jobExperience"`
	PreviousJobs  []string `json:"previousJobs"`
	Town          string   `json:"town"`
	NameJob       string   `json:"nameJob"`
	Age           int      `json:"age"`
}

func NewVacancy(nameFirst string, secondName string, jobExp int, previousJobs []string, towner string, nameJob string, age int, id string) *VacancyStruct {
	return &VacancyStruct{
		ID:            id,
		FirstName:     nameFirst,
		SecondName:    secondName,
		JobExperience: jobExp,
		PreviousJobs:  previousJobs,
		Town:          towner,
		NameJob:       nameJob,
		Age:           age,
	}
}

func (v *VacancyStruct) PrintVacancy() {
	fmt.Println(v)
}
