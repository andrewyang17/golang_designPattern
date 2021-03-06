package main

type stationManager struct {
	isPlatformFree bool
	trainQueue     []train
}

func (s *stationManager) canArrive(t train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *stationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		dequeueTrain := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		dequeueTrain.permitArrival()
	}
}

func NewStationManager() *stationManager {
	return &stationManager{
		isPlatformFree: true,
	}
}
