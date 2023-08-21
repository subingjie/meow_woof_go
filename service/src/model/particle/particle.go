package particle

type Particle struct {
	Id         int64  `json:"id"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Disable    int    `json:"disable"`

	Name    string `json:"name"`
	Content string `json:"content"`
}
