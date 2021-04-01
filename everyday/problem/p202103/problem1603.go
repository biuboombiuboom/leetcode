package p202103

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor2(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big, medium, small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		this.big--
		return this.big >= 0
	case 2:
		this.medium--
		return this.medium >= 0
	case 3:
		this.small--
		return this.small >= 0

	}
	return false
}
