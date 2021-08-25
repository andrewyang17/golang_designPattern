package main

type player struct {
	dress      dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{
		dress:      dress,
		playerType: playerType,
	}
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}