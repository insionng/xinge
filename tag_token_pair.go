package xinge

type TagTokenPair struct {
	Tag   string `json:"tag,omitempty"`
	Token string `json:"token,omitempty"`
}

func NewTagTokenPair(tag, token string) *TagTokenPair {
	return &TagTokenPair{tag, token}
}
