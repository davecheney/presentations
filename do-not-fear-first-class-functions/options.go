package world

func WithReticulatedSplines(c *Config) {}

type Terrain struct {
	config Config
}

func NewTerrain(opts ...func(*Config)) *Terrain {
	var t Terrain
	for _, fn := range opts {
		fn(&t.Config)
	}
	return &t
}

func main() {
	t := NewTerrain(WithReticulatedSplines)
	// ....
}
