package structs

type Reponse struct {
	Data string `json:"data"`
	Tag  string `json:"tag,omitempty"`
}
