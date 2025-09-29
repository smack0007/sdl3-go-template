package main

import (
	"math"
	"os"
	"runtime"

	"github.com/smack0007/sdl3-go/sdl"
)

const WINDOW_WIDTH = 1024
const WINDOW_HEIGHT = 768
const WINDOW_TITLE = "App"
const DESIRED_FPS = 60

var TICK_RATE = (uint64)(math.Floor(float64(1000.0) / (float64)(DESIRED_FPS)))

func main() {
	runtime.LockOSThread()
	os.Exit(run())
}

func run() int {
	err := sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "Failed initialize SDL.")
		return 1
	}

	defer sdl.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(WINDOW_TITLE, WINDOW_WIDTH, WINDOW_HEIGHT, sdl.WINDOW_OCCLUDED)

	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "Failed to create window and renderer.")
		return 1
	}

	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	shouldQuit := false
	var event sdl.Event

	currentTime := sdl.GetTicks()
	lastTime := currentTime

	for !shouldQuit {
		for sdl.PollEvent(&event) {
			switch event.Type() {

			case sdl.EVENT_QUIT:
				shouldQuit = true
			}
		}

		currentTime = sdl.GetTicks()
		elapsedTime := currentTime - lastTime

		if elapsedTime >= TICK_RATE {
			update(float32(elapsedTime) / float32(1000))
			draw(renderer)

			lastTime = currentTime
		}

		sdl.Delay(1)
	}

	return 0
}

func update(elapsedTime float32) {
}

func draw(renderer *sdl.Renderer) {
	sdl.SetRenderDrawColor(renderer, 100, 149, 237, 255)
	sdl.RenderClear(renderer)

	sdl.RenderPresent(renderer)
}
