package track

import "github.com/bars-p/omp-bot/internal/model/logistic"


type TrackService interface {
	Describe(trackID uint64) (*logistic.Track, error)
	List(cursor, limit uint64) ([]logistic.Track, error)
	Create(logistic.Track) (uint64, error)
	Update(trackID uint64, track logistic.Track) error
	Remove(trackID uint64) (bool, error)
}


type DummyTrackService struct{}

func NewDummyTrackService() *DummyTrackService {
	return &DummyTrackService{}
}

func (dts *DummyTrackService) Describe(idx int) (*logistic.Track, error) {
	return &allEntities[idx], nil
}

func (dts *DummyTrackService) List() []logistic.Track {
	return allEntities
}
