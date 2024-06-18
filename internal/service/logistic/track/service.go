package track

import "github.com/bars-p/omp-bot/internal/model/logistic"


type TrackService interface {
	Describe(trackID uint64) (*logistic.Track, error)
	Get(idx int) (*logistic.Track, error)
	List(cursor, limit uint64) ([]logistic.Track, error)
	Create(logistic.Track) (uint64, error)
	Update(trackID uint64, track logistic.Track) error
	Remove(trackID uint64) (bool, error)
}


type DummyTrackService struct{}

func NewDummyTrackService() *DummyTrackService {
	return &DummyTrackService{}
}

func (dts *DummyTrackService) Describe(idx uint64) (*logistic.Track, error) {
	return &allEntities[0], nil
}

func (dts *DummyTrackService) Get(idx int) (*logistic.Track, error) {
	return &allEntities[idx], nil
}

func (dts *DummyTrackService) List(cursor, limit uint64) ([]logistic.Track, error) {
	return allEntities, nil
}

func (dts *DummyTrackService) Create(logistic.Track) (uint64, error) {
	return 0, nil
}

func (dts *DummyTrackService) Update(trackID uint64, track logistic.Track) error {
	return nil
}

func (dts *DummyTrackService) Remove(trackID uint64) (bool, error) {
	return true, nil
}
