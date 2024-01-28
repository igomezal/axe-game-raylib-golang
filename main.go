package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type point struct {
	x int32
	y int32
}

type bounds struct {
	left_x  int32
	right_x int32
	upper_y int32
	lower_y int32
}

type entity struct {
	position   point
	bounderies bounds
	width      int32
}

type player struct {
	entity
}

type enemy struct {
	entity
}

func NewPlayer(origin point, width int32) player {
	return player{
		entity{
			origin,
			bounds{
				left_x:  origin.x - width,
				right_x: origin.x + width,
				upper_y: origin.y - width,
				lower_y: origin.y + width,
			},
			width,
		},
	}
}

func NewAxe(origin point, width int32) enemy {
	return enemy{
		entity{
			origin,
			bounds{
				left_x:  origin.x,
				right_x: origin.x + width,
				upper_y: origin.y,
				lower_y: origin.y + width,
			},
			width,
		},
	}
}

func (e *entity) Move(directionX, directionY int32) {
	e.position.x += directionX
	e.position.y += directionY
}

func (p *player) UpdateBounderies() {
	p.bounderies.left_x = p.position.x - p.width
	p.bounderies.right_x = p.position.x + p.width
	p.bounderies.upper_y = p.position.y - p.width
	p.bounderies.lower_y = p.position.y + p.width
}

func (e *enemy) UpdateBounderies() {
	e.bounderies.left_x = e.position.x
	e.bounderies.right_x = e.position.x + e.width
	e.bounderies.upper_y = e.position.y
	e.bounderies.lower_y = e.position.y + e.width
}

func (a *entity) HaveCollided(b entity) bool {
	return (a.bounderies.lower_y >= b.bounderies.upper_y) &&
		(a.bounderies.upper_y <= b.bounderies.lower_y) &&
		(a.bounderies.left_x <= b.bounderies.right_x) &&
		(a.bounderies.right_x >= b.bounderies.left_x)
}

func main() {
	var width int32 = 800
	var height int32 = 450
	var direction int32 = 10

	rl.InitWindow(width, height, "Hello Window!")

	rl.SetTargetFPS(60)

	var player player = NewPlayer(point{width / 2, height / 2}, 25)
	var axe enemy = NewAxe(point{300, 0}, 50)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		if player.HaveCollided(axe.entity) {
			rl.DrawText("Game Over!", width/2, height/2, 20, rl.Red)
		} else {
			rl.DrawCircle(player.position.x, player.position.y, float32(player.width), rl.Blue)
			rl.DrawRectangle(axe.position.x, axe.position.y, axe.width, axe.width, rl.Red)

			axe.Move(0, direction)
			axe.UpdateBounderies()
			if axe.position.y+axe.width > height || axe.position.y < 0 {
				direction = -direction
			}

			if rl.IsKeyDown(rl.KeyD) && player.position.x < width-player.width {
				player.Move(10, 0)
			}
			if rl.IsKeyDown(rl.KeyA) && player.position.x > 0+player.width {
				player.Move(-10, 0)
			}

			player.UpdateBounderies()
		}

		rl.EndDrawing()
	}
}
