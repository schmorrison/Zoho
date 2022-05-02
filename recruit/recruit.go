package recruit

import (
	"math/rand"
	"time"

	zoho "github.com/schmorrison/Zoho"
)

type Module string

const (
	CandidatesModule   Module = "Candidates"
	ClientsModule      Module = "Clients"
	ContactsModule     Module = "Contacts"
	CustomModule1      Module = "CustomModule1"
	CustomModule2      Module = "CustomModule2"
	CustomModule3      Module = "CustomModule3"
	CustomModule4      Module = "CustomModule4"
	CustomModule5      Module = "CustomModule5"
	InterviewsModule   Module = "Interviews"
	JobOpeningsModule  Module = "JobOpenings"  // this is used mostly
	Job_OpeningsModule Module = "Job_Openings" // inconsistency in Zoho API
)

type RelatedRecord string

const (
	Activities          RelatedRecord = "Activities"
	Activity_History    RelatedRecord = "Activity_History"
	Attachments         RelatedRecord = "Attachments"
	Campaigns           RelatedRecord = "Campaigns"
	Contacts            RelatedRecord = "Contacts"
	Custom_Related_List RelatedRecord = "Custom_Related_List"
	Interviews          RelatedRecord = "Interviews"
	Invited_Events      RelatedRecord = "Invited_Events"
	Notes               RelatedRecord = "Notes"
	Offers              RelatedRecord = "Offers"
)

type AttachmentCategory string

const (
	Contracts        AttachmentCategory = "Contracts"
	Cover_Letter     AttachmentCategory = "Cover_Letter"
	Formatted_Resume AttachmentCategory = "Formatted_Resume"
	Offer            AttachmentCategory = "Offer"
	Others           AttachmentCategory = "Others"
	Resume           AttachmentCategory = "Resume"
)

// API is used for interacting with the Zoho recruit API
// the exposed methods are primarily access to recruit modules which provide access to recruit Methods
type API struct {
	*zoho.Zoho
	id string
}

// New returns a *recruit.API with the provided zoho.Zoho as an embedded field
func New(z *zoho.Zoho) *API {
	id := func() string {
		var id []byte
		keyspace := "abcdefghijklmnopqrutuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		for i := 0; i < 25; i++ {
			source := rand.NewSource(time.Now().UnixNano())
			rnd := rand.New(source)
			id = append(id, keyspace[rnd.Intn(len(keyspace))])
		}
		return string(id)
	}()

	return &API{
		Zoho: z,
		id:   id,
	}
}
