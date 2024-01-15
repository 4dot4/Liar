package main

func intro(r *rayLogo, screen *screens) {
	if r.state == 0 { // State 0: Small box blinking
		r.framesCounter++

		if r.framesCounter == 120 {
			r.state = 1
			r.framesCounter = 0 // Reset counter... will be used later...
		}
	} else if r.state == 1 { // State 1: Top and left bars growing
		r.topSideRecWidth += 4
		r.leftSideRecHeight += 4

		if r.topSideRecWidth == 256 {
			r.state = 2
		}
	} else if r.state == 2 { // State 2: Bottom and right bars growing
		r.bottomSideRecWidth += 4
		r.rightSideRecHeight += 4

		if r.bottomSideRecWidth == 256 {
			r.state = 3
		}
	} else if r.state == 3 { // State 3: Letters appearing (one by one)
		r.framesCounter++

		if r.framesCounter%12 == 0 { // Every 12 frames, one more letter!
			r.lettersCount++
			r.framesCounter = 0
		}

		if r.lettersCount >= 6 { // When all letters have appeared, just fade out everything
			r.alpha -= 0.02

			if r.alpha <= 0.0 {
				r.alpha = 0.0
				*screen = Client
			}
		}
	}

}
func update(game *game) {
	switch game.GameScreen {
	case Start:
		intro(&game.RayLogo, &game.GameScreen)
	}
}
