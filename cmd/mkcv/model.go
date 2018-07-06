package main

// https://jsonresume.org/schema/
type (
	CV struct {
		Headings     Headings      `json:"headings"`
		Basics       Basics        `json:"basics"`
		Education    []Education   `json:"education"`
		Work         []Work        `json:"work"`
		Skills       []Skill       `json:"skills"`
		Projects     []Project     `json:"projects"`
		Awards       []Award       `json:"awards"`
		Sections     []string      `json:"sections"`
		Volunteer    []Volunteer   `json:"volunteer"`
		Publications []Publication `json:"publications"`
		Languages    []Language    `json:"languages"`
		Interests    []Interest    `json:"interests"`
		References   []Reference   `json:"references"`
	}
	Reference struct {
		Name      string `json:"name"`
		Reference string `json:"reference"`
	}
	Language struct {
		Name  string `json:"name"`
		Level string `json:"level"`
	}
	Interest struct {
		Name     string   `json:"name"`
		Keywords []string `json:"keywords"`
	}
	Publication struct {
		Name        string `json:"name"`
		Publisher   string `json:"publisher"`
		ReleaseDate string `json:"releaseDate"`
		Website     string `json:"website"`
		Summary     string `json:"summary"`
	}
	Volunteer struct {
		Organization string   `json:"organization"`
		Position     string   `json:"position"`
		Website      string   `json:"website"`
		StartDate    string   `json:"startDate"`
		EndDate      string   `json:"endDate"`
		Summary      string   `json:"summary"`
		Highlights   []string `json:"highlights"`
	}
	Headings struct {
		Projects  string `json:"projects"`
		Awards    string `json:"awards"`
		Education string `json:"education"`
		Work      string `json:"work"`
		Skills    string `json:"skills"`
	}
	Basics struct {
		Name     string    `json:"name"`
		Label    string    `json:"label"`
		Picture  string    `json:"picture"`
		Email    string    `json:"email"`
		Phone    string    `json:"phone"`
		Website  string    `json:"website"`
		Linkedin string    `json:"linkedin"`
		Summary  string    `json:"summary"`
		Location Location  `json:"location"`
		Profiles []Profile `json:"profiles"`
	}
	Location struct {
		Address     string `json:"address"`
		PostalCode  string `json:"postalCode"`
		City        string `json:"city"`
		CountryCode string `json:"countryCode"`
		Region      string `json:"region"`
	}
	Profile struct {
		Network  string `json:"network"`
		Username string `json:"username"`
		URL      string `json:"url"`
	}
	Education struct {
		Institution string   `json:"institution"`
		Location    string   `json:"location"`
		Area        string   `json:"area"`
		StudyType   string   `json:"studyType"`
		StartDate   string   `json:"startDate"`
		EndDate     string   `json:"endDate"`
		GPA         string   `json:"gpa"`
		Courses     []string `json:"courses"`
	}
	Work struct {
		Position   string   `json:"position"`
		Company    string   `json:"company"`
		Location   string   `json:"location"`
		Website    string   `json:"website,omitempty"`
		StartDate  string   `json:"startDate"`
		EndDate    string   `json:"endDate"`
		Highlights []string `json:"highlights"`
	}
	Skill struct {
		Name     string   `json:"name"`
		Level    string   `json:"level"`
		Keywords []string `json:"keywords"`
	}
	Project struct {
		Name       string   `json:"name"`
		Summary    string   `json:"summary"`
		StartDate  string   `json:"startDate"`
		EndDate    string   `json:"endDate"`
		URL        string   `json:"url"`
		Keywords   []string `json:"keywords"`
		Highlights []string `json:"highlights"`
	}
	Award struct {
		Title   string `json:"title"`
		Date    string `json:"date"`
		Awarder string `json:"awarder"`
		Summary string `json:"summary"`
	}
)
