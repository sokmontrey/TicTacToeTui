package vec2

type Vec2 struct {
	X int
	Y int
}

func NewVec2(x int, y int) Vec2 {
	return Vec2{X: x, Y: y}
}

func ZeroVec2() Vec2 {
	return Vec2{X: 0, Y: 0}
}

func (v Vec2) To(x int, y int) Vec2 {
	return Vec2{X: x, Y: y}
}

func (v Vec2) Up() Vec2 {
	return Vec2{X: v.X, Y: v.Y - 1}
}

func (v Vec2) Down() Vec2 {
	return Vec2{X: v.X, Y: v.Y + 1}
}

func (v Vec2) Left() Vec2 {
	return Vec2{X: v.X - 1, Y: v.Y}
}

func (v Vec2) Right() Vec2 {
	return Vec2{X: v.X + 1, Y: v.Y}
}

func (v Vec2) Subtract(other Vec2) Vec2 {
	return Vec2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}
