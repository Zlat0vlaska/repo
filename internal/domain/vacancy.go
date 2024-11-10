package domain

type VacancyState string

const (
	VacancyStateOpen    VacancyState = "open"
	VacancyStateClosed  VacancyState = "closed"
	VacancyStateArchive VacancyState = "archive"
)

type Vacancy struct {
	ID          string   `db:"id"`
	Name        string   `db:"name"`
	Company     string   `db:"company"`
	State       string   `db:"state"`
	JobTitle    string   `db:"job_title"`
	Salary      string   `db:"salary"`
	Description string   `db:"description"`
	Grades      []string `db:"grades"`
	DateCreate  string   `db:"date_create"`
	IsFavorite  bool     `db:"is_favorite"`
	Skills      []string `db:"skills"`
	Country     string   `db:"country"`
	Regions     []string `db:"regions"`
	Cities      []string `db:"cities"`
}

type VacancyFilter struct {
	Keyword   string
	Country   string
	Region    string
	City      string
	Grade     string
	SortBy    string
	HrID      string
	SalaryMin int
	SalaryMax int
	Limit     int
	Offset    int
}
