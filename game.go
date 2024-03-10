package main

func gameCheckCollision() (bool, string) {
	switch playfieldArray[playerYpos][playerXpos] {
	case 1:
		return true, "Wall 1"
	case 2:
		return true, "Wall 2"
	case 3:
		return true, "Wall 3"
	case 4:
		return false, "Water 1"
	case 5:
		return true, "Obstacle 1"
	case 6:
		return true, "Obstacle 2"
	case 7:
		return true, "Bonus 1"
	case 8:
		return true, "Bonus 2"
	case 9:
		return true, "Finish"
	}
	return false, "None"
}
