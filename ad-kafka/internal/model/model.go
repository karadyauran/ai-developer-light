package model

type Request struct {
	ID        string `json:"ID"`
	Params    Params `json:"params"`
	Topic     string `json:"topic"`
	CreatedAt string `json:"created_at"`
}

type Params struct {
	ProjectType    string `json:"project_type"`
	Language       string `json:"language"`
	GenerationType string `json:"generation_type"` // random idea or custom
	Idea           string `json:"idea"`            // if needed
	Token          string `json:"token"`           // user github token
}
