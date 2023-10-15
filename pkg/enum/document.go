package enum

import (
	"fmt"
	"io"
	"strconv"
)

type Dir string

const (
	DirEmailTemplates Dir = "mail-templates"
	DirPdfTemplates   Dir = "pdf-templates"
	DirSecrets        Dir = "secret"
	DirErrorLogs      Dir = "error-logs"
)

func (t Dir) String() string {
	return string(t)
}

type DocFolder string

const (
	FolderPartnerDocs    DocFolder = "PartnerDocs"
	FolderJobDocs        DocFolder = "JobDocs"
	FolderEstimates      DocFolder = "Estimates"
	FolderTrainingVideos DocFolder = "TrainingVideos"
	FolderPublicData     DocFolder = "PublicData"
	FolderSurvey         DocFolder = "Survey"
)

var AllDocFolder = []DocFolder{
	FolderPartnerDocs,
	FolderJobDocs,
	FolderEstimates,
	FolderTrainingVideos,
	FolderPublicData,
	FolderSurvey,
}

func (t DocFolder) String() string {
	return string(t)
}

// Values provides list valid values for Enum.
func (DocFolder) Values() (kinds []string) {
	for _, s := range AllDocFolder {
		kinds = append(kinds, string(s))
	}
	return
}

func (t DocFolder) IsValid() bool {
	for _, t := range AllDocFolder {
		if t == t {
			return true
		}
	}
	return false
}

func (t *DocFolder) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*t = DocFolder(str)
	if !t.IsValid() {
		return fmt.Errorf("%s is not a valid DocFolder", str)
	}
	return nil
}

func (t DocFolder) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(t.String()))
}

// DocSection type
type DocSection string

const (
	SectionDocs           DocSection = "Docs"
	SectionInspectionDocs DocSection = "InspectionDocs"
	SectionProductionPics DocSection = "ProductionPics"
	SectionPricingPDF     DocSection = "PricingPDF"
	SectionProof          DocSection = "Proof"
	SectionLogo           DocSection = "Logo"
	SectionAvatar         DocSection = "Avatar"
	SectionImage          DocSection = "Image"
	SectionVideo          DocSection = "Video"
	SectionProposal       DocSection = "Proposal"
)

var AllDocSection = []DocSection{
	SectionDocs,
	SectionInspectionDocs,
	SectionProductionPics,
	SectionPricingPDF,
	SectionProof,
	SectionLogo,
	SectionAvatar,
	SectionImage,
	SectionVideo,
	SectionProposal,
}

func (t DocSection) String() string {
	return string(t)
}

// Values provides list valid values for Enum.
func (DocSection) Values() (kinds []string) {
	for _, s := range AllDocSection {
		kinds = append(kinds, string(s))
	}
	return
}

func (t DocSection) IsValid() bool {
	for _, t := range AllDocSection {
		if t == t {
			return true
		}
	}
	return false
}

func (t *DocSection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*t = DocSection(str)
	if !t.IsValid() {
		return fmt.Errorf("%s is not a valid DocSection", str)
	}
	return nil
}

func (t DocSection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(t.String()))
}
