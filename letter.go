package cv

type (
	CoverLetter struct {
		Job Job `json:"job"`
		CV  CV  `json:"cv"`
	}
	Job struct {
		Company   Company   `json:"company"`
		Position  string    `json:"position"`
		Recruiter Recruiter `json:"recruiter"`
		Channel   string    `json:"channel"`
	}
	Company struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	Recruiter struct {
		Name  string `json:"name"`
		Title string `json:"title"`
	}
)
