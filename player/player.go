package player

import rl "github.com/gen2brain/raylib-go/raylib"

type player struct {
	x, y   int
	health int
	speed  int
}

func NewPlayer(x, y, speed int, health int) *player {
	return &player{
		x:      x,
		y:      y,
		health: health,
		speed:  speed,
	}
}

func (p *player) Move() {
	if rl.IsKeyDown(rl.KeyW) {
		p.y -= p.speed
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.y += p.speed
	}
	if rl.IsKeyDown(rl.KeyA) {
		p.x -= p.speed
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.x += p.speed
	}
}

func (p *player) Render() {
	rl.DrawCircle(int32(p.x), int32(p.y), 32, rl.Red)
}
