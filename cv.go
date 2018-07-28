package cv

import "fmt"

// https://jsonresume.org/schema/
type (
	CV struct {
		Title        string        `json:"title"`
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
		Hidden    bool   `json:"hidden"`
	}
	Name struct {
		First  string `json:"first"`
		Nick   string `json:"nick"`
		Second string `json:"second"`
	}
	Language struct {
		Name   string `json:"name"`
		Level  string `json:"level"`
		Hidden bool   `json:"hidden"`
	}
	Interest struct {
		Name     string   `json:"name"`
		Keywords []string `json:"keywords"`
		Hidden   bool     `json:"hidden"`
	}
	Publication struct {
		Name        string `json:"name"`
		Publisher   string `json:"publisher"`
		ReleaseDate string `json:"releaseDate"`
		Website     string `json:"website"`
		Summary     string `json:"summary"`
		Hidden      bool   `json:"hidden"`
	}
	Volunteer struct {
		Organization string   `json:"organization"`
		Position     string   `json:"position"`
		Website      string   `json:"website"`
		StartDate    string   `json:"startDate"`
		EndDate      string   `json:"endDate"`
		Summary      string   `json:"summary"`
		Highlights   []string `json:"highlights"`
		Hidden       bool     `json:"hidden"`
	}
	Headings struct {
		Projects  string `json:"projects"`
		Awards    string `json:"awards"`
		Education string `json:"education"`
		Work      string `json:"work"`
		Skills    string `json:"skills"`
	}
	Basics struct {
		Name     Name      `json:"name"`
		Label    string    `json:"label"`
		Picture  string    `json:"picture"`
		Email    string    `json:"email"`
		Phone    string    `json:"phone"`
		Website  string    `json:"website"`
		Linkedin string    `json:"linkedin"`
		Summary  string    `json:"summary"`
		Skype    Skype     `json:"skype"`
		Github   string    `json:"github"`
		Location Location  `json:"location"`
		Profiles []Profile `json:"profiles"`
	}
	Skype struct {
		Email string `json:"email"`
		ID    string `json:"id"`
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
		Highlights  []string `json:"highlights"`
		Hidden      bool     `json:"hidden"`
	}
	Work struct {
		Position   string   `json:"position"`
		Company    string   `json:"company"`
		Location   string   `json:"location"`
		Website    string   `json:"website,omitempty"`
		StartDate  string   `json:"startDate"`
		EndDate    string   `json:"endDate"`
		Highlights []string `json:"highlights"`
		Hidden     bool     `json:"hidden"`
	}
	Skill struct {
		Name     string   `json:"name"`
		Level    string   `json:"level"`
		Keywords []string `json:"keywords"`
		Hidden   bool     `json:"hidden"`
	}
	Project struct {
		Name       string   `json:"name"`
		Summary    string   `json:"summary"`
		StartDate  string   `json:"startDate"`
		EndDate    string   `json:"endDate"`
		URL        string   `json:"url"`
		Keywords   []string `json:"keywords"`
		Highlights []string `json:"highlights"`
		Hidden     bool     `json:"hidden"`
	}
	Award struct {
		Title   string `json:"title"`
		Date    string `json:"date"`
		Awarder string `json:"awarder"`
		Summary string `json:"summary"`
		Hidden  bool   `json:"hidden"`
	}
)

func (n Name) String() string {
	if n.Nick != "" {
		return fmt.Sprintf(`%s "%s" %s`, n.First, n.Nick, n.Second)
	}
	return n.First + " " + n.Second
}
