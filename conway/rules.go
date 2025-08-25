package conway

func checkLiveCell(neighbours uint8) bool {
	switch neighbours {
	case 2, 3:
		return true
	default:
		return false
	}
}

func checkDeadCell(neighbours uint8) bool {
	if neighbours == 3 {
		return true
	} else {
		return false
	}
}
