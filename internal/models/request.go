package models

type pastExperiences struct {
	Sales   bool `json:"sales,omitempty"`
	Support bool `json:"support,omitempty"`
}

type internetTest struct {
	DownloadSpeed float32 `json:"download_speed,omitempty"`
	UploadSpeed   float32 `json:"upload_speed,omitempty"`
}

type Request struct {
	Age             uint8            `json:"age,omitempty"`
	EducationLevel  string           `json:"education_level,omitempty"`
	PastExperiences *pastExperiences `json:"past_experiences,omitempty"`
	InternetTest    *internetTest    `json:"internet_test,omitempty"`
	WritingScore    float32          `json:"writing_score,omitempty"`
	ReferralCode    string           `json:"referral_code,omitempty"`
}
