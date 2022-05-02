package recruit

import "time"

type InsertCandidatesData struct {
	Origin                   string    `json:"Origin"`
	Email                    string    `json:"Email"`
	CurrencySymbol           string    `json:"$currency_symbol"`
	LastActivityTime         time.Time `json:"Last_Activity_Time"`
	HighestQualificationHeld string    `json:"Highest_Qualification_Held"`
	SkillSet                 string    `json:"Skill_Set"`
	Converted                bool      `json:"$converted"`
	ProcessFlow              bool      `json:"$process_flow"`
	UpdatedOn                time.Time `json:"Updated_On"`
	CurrentEmployer          string    `json:"Current_Employer"`
	Street                   string    `json:"Street"`
	ZipCode                  string    `json:"Zip_Code"`
	ID                       string    `json:"id"`
	ExperienceInYears        int       `json:"Experience_in_Years"`
	Approved                 bool      `json:"$approved"`
	Approval                 struct {
		Delegate bool `json:"delegate"`
		Approve  bool `json:"approve"`
		Reject   bool `json:"reject"`
		Resubmit bool `json:"resubmit"`
	} `json:"$approval"`
	CandidateStatus string      `json:"Candidate_Status"`
	CandidateID     string      `json:"Candidate_ID"`
	LastMailedTime  interface{} `json:"Last_Mailed_Time"`
	CreatedTime     time.Time   `json:"Created_Time"`
	Followed        bool        `json:"$followed"`
	CandidateOwner  struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"Candidate_Owner"`
	Editable       bool          `json:"$editable"`
	IsLocked       bool          `json:"Is_Locked"`
	City           string        `json:"City"`
	IsUnqualified  bool          `json:"Is_Unqualified"`
	AssociatedTags []interface{} `json:"Associated_Tags"`
	AdditionalInfo string        `json:"Additional_Info"`
	State          string        `json:"State"`
	Country        string        `json:"Country"`
	CreatedBy      struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"Created_By"`
	SecondaryEmail      string      `json:"Secondary_Email"`
	IsAttachmentPresent bool        `json:"Is_Attachment_Present"`
	Rating              int         `json:"Rating"`
	AppliedWithLinkedin interface{} `json:"$applied_with_linkedin"`
	Website             string      `json:"Website"`
	Twitter             string      `json:"Twitter"`
	CurrentJobTitle     string      `json:"Current_Job_Title"`
	Salutation          string      `json:"Salutation"`
	Source              string      `json:"Source"`
	FirstName           string      `json:"First_Name"`
	FullName            string      `json:"Full_Name"`
	ModifiedBy          struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"Modified_By"`
	SkypeID            string        `json:"Skype_ID"`
	ExperienceDetails  []interface{} `json:"Experience_Details"`
	Phone              string        `json:"Phone"`
	EmailOptOut        bool          `json:"Email_Opt_Out"`
	EducationalDetails []interface{} `json:"Educational_Details"`
	ConvertedDetail    struct {
	} `json:"$converted_detail"`
	CareerPageInviteStatus      string      `json:"Career_Page_Invite_Status"`
	Mobile                      string      `json:"Mobile"`
	LastName                    string      `json:"Last_Name"`
	CurrentSalary               interface{} `json:"Current_Salary"`
	AssociatedAnySocialProfiles bool        `json:"Associated_any_Social_Profiles"`
	Fax                         string      `json:"Fax"`
	ExpectedSalary              interface{} `json:"Expected_Salary"`
}
