package world

func WithReticulatedSplines(c *Config)

type Terrain struct {
	config Config
}

func NewTerrain(options ...func(*Config)) *Terrain {
	var t Terrain
	for _, option := range options {
		option(&t.Config)
	}
	return &t
}

func main() {
	t := NewTerrain(WithReticulatedSplines)
	// ....
}
