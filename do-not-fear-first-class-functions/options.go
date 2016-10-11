package world

func WithReticulatedSplines(c *Config)

type Config struct{}

type Terrain struct {
	config Config
}

func NewTerrain(options ...func(*Config)) *Terrain {
	var t Terrain
	for _, option := range options {
		option(&t.config)
	}
	return &t
}

func main() {
	t := NewTerrain(WithReticulatedSplines)
	// [ simulation intensifies ]
}
