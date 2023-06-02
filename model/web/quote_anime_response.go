package web

type QuoteAnimeResponse struct {
	Anime         string `json:"anime,omitempty"`
	CharacterName string `json:"character_name,omitempty"`
	Quote         string `json:"quote,omitempty"`
}
