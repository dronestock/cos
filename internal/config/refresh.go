package config

type Refresh struct {
	Url   string   `json:"url,omitempty" validate:"required_without=Urls"`
	Urls  []string `json:"urls,omitempty" validate:"required_without=Url"`
	Path  string   `json:"path,omitempty" validate:"required_without=Paths"`
	Paths []string `json:"paths,omitempty" validate:"required_without=Path"`
	Regin string   `default:"ap-chengdu" json:"regin,omitempty"`
	Type  string   `default:"delete" json:"type,omitempty" validate:"oneof=flush delete"`
}
