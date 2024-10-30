package model

type DockerContainer struct {
	Hostname string   `json:"hostname"`
	Image    string   `json:"image"`
	Files    []File   `json:"files"`
	Commands []string `json:"commands"`
}

type File struct {
	Filename string `json:"filename"`
}
