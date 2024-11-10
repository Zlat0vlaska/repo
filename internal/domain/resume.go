package domain

type Resume struct {
	ID                     string   `db:"id"`
	CandidateName          string   `db:"candidate_name"`
	Status                 string   `db:"status"`
	Grade                  string   `db:"grade"`
	Country                string   `db:"country"`
	City                   string   `db:"city"`
	Profession             string   `db:"profession"`
	Schedule               string   `db:"schedule"`
	Citizenship            string   `db:"citizenship"`
	BusinessTripsReadiness string   `db:"business_trips_readiness"`
	Permission             string   `db:"permission"`
	Salary                 float64  `db:"salary"`
	RelocationReadiness    string   `db:"relocation_readiness"`
	ExperienceYears        int      `db:"experience_years"`
	Skills                 []string `db:"skills"`
	Languages              []string `db:"languages"`
	Education              string   `db:"education"`
}
type ResumeFilter struct {
	Keyword                string
	Profession             string
	Schedule               string
	Citizenship            string
	BusinessTripsReadiness string
	SalaryMin              int
	SalaryMax              int
	RelocationReadiness    string
	Skills                 string
	Languages              string
	Education              string
	Limit                  int
	Offset                 int
	SortBy                 string
	SortOrder              string
}
