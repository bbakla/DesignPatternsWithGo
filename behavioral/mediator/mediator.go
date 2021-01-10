package mediator

type Mediator interface {
	canLand(Airplane) bool
	notifyFree()
}

type AirTower struct {
	airplanes           []Airplane
	NumberOfRunways     int
	numberOfUsedRunways int
}

func NewAirTower(numberOfRunways int) Mediator {

	return &AirTower{
		NumberOfRunways: numberOfRunways,
	}
}

func (tower *AirTower) canLand(airplane Airplane) bool {
	if tower.NumberOfRunways > tower.numberOfUsedRunways {
		airplane.PermitArrival()
		tower.numberOfUsedRunways++

		return true
	} else {
		tower.airplanes = append(tower.airplanes, airplane)

		return false
	}
}

func (tower *AirTower) notifyFree() {
	tower.numberOfUsedRunways--

	if len(tower.airplanes) > 0 {
		waitingAirplane := tower.airplanes[0]
		waitingAirplane.PermitArrival()

		tower.airplanes = tower.airplanes[1:]
	}

}
