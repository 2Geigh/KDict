package models

type TemplateData struct {
	SearchQuery   string
	SearchResults []DictSearch
}

type DictSearch struct {
	Title   string     `xml:"title"`
	Total   int        `xml:"total"`
	Results []DictItem `xml:"item"`
}

type DictItem struct {
	Target_code      int            `xml:"target_code"`
	Word             string         `xml:"word"`
	Sup_no           int            `xml:"sup_no"`
	Etymology        string         `xml:"origin"`
	Pronunciation    string         `xml:"pronunciation"`
	Word_grade_level string         `xml:"word_grade"`
	Word_type        string         `xml:"pos"`
	Entry_link       string         `xml:"link"`
	Sense            DictEntrySense `xml:"sense"`
}

type DictEntrySense struct {
	Order      int    `xml:"sense_order"`
	Definition string `xml:"definition"`
}
