package config

type Refresh struct {
	Url   string   `json:"url,omitempty" validate:"required_without_all=Urls Path Paths"`
	Urls  []string `json:"urls,omitempty" validate:"required_without_all=Url Path Paths"`
	Path  string   `json:"path,omitempty" validate:"required_without_all=Paths Url Urls"`
	Paths []string `json:"paths,omitempty" validate:"required_without_all=Path Url Urls"`
	Regin string   `default:"ap-chengdu" json:"regin,omitempty"`
	Type  string   `default:"delete" json:"type,omitempty" validate:"oneof=flush delete"`
}
