package model

type Filters struct {
	Year       uint64
	ComedianId uint64
	SortBy     string // popular || new
	ExcludedId uint64
	Limit      uint64
}

type YoutubeDownloadBody struct {
	Link     string `json:"link"`
	OutDir   string `json:"outDir"`
	Filename string `json:"filename"`
}

type UpdateConcertDto struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
