package models

type Book struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Feature     string `json:"feature"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Remarks     string `json:"remarks"`
	Lat         string `json:"lat"`
	Long        string `json:"long"`
	// Lattitude   float64 `json:"lat"`
	// Longitude   float64 `json:"long"`
	Tags  string `json:"tags"`
	Tags1 string `json:"tags1"`
	Tags2 string `json:"tags2"`
	// Tags        map[string]interface{} `json:"tags" sql:"type:jsonb"`
}
