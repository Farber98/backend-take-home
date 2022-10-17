package models

type PastExperiences struct {
	Sales   bool `json:"sales,omitempty"`
	Support bool `json:"support,omitempty"`
}

type InternetTest struct {
	DownloadSpeed float32 `json:"download_speed,omitempty"`
	UploadSpeed   float32 `json:"upload_speed,omitempty"`
}

type Request struct {
	Age             uint8            `json:"age,omitempty"`
	EducationLevel  string           `json:"education_level,omitempty"`
	PastExperiences *PastExperiences `json:"past_experiences,omitempty"`
	InternetTest    *InternetTest    `json:"internet_test,omitempty"`
	WritingScore    float32          `json:"writing_score,omitempty"`
	ReferralCode    string           `json:"referral_code,omitempty"`
}
