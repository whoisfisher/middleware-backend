package entity

type KubesphereUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}
