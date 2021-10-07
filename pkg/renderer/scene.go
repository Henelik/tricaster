package renderer

type Scene struct {
	Name   string
	Camera *Camera
	World  *World
}

func NewScene(config *Configuration) *Scene {
	world := config.World.ToWorld()

	world.Geometry = make([]Primitive, 0, len(config.Objects))

	for _, object := range config.Objects {
		world.Geometry = append(world.Geometry, object.ToPrimitive())
	}

	return &Scene{
		Name:   config.Name,
		World:  world,
		Camera: config.Camera.ToCamera(),
	}
}

func (s *Scene) Render() {
	s.Camera.GoRender(s.World).SaveImage(s.Name + ".png")
}
