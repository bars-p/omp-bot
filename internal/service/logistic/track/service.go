package track

import (
	"fmt"

	"github.com/bars-p/omp-bot/internal/model/logistic"
)


type SubdomainService interface {
	Describe(trackID uint64) (*logistic.Track, error)
	Get(idx int) (*logistic.Track, error)
	List(cursor, limit uint64) ([]logistic.Track, error)
	Create(title string) (uint64, error)
	Update(trackID uint64, track logistic.Track) error
	Remove(trackID uint64) (bool, error)
}


type TrackService struct{
	LastID uint64
}

func NewTrackService() *TrackService {
	return &TrackService{ LastID: uint64(len(allEntities)+1) }
}

func (ts *TrackService) Describe(trackID uint64) (*logistic.Track, error) {
	return &allEntities[0], nil
}

func (ts *TrackService) Get(idx int) (*logistic.Track, error) {
	l := len(allEntities)
	if  l == 0 || idx >= l || idx < 0 {
		return nil, fmt.Errorf("index %d is out of range", idx)
	}
	return &allEntities[idx], nil
}

func (ts *TrackService) List(cursor, limit uint64) ([]logistic.Track, error) {
	l := len(allEntities)
	if cursor >= uint64(l) {
		return nil, fmt.Errorf("cursor out of range")
	} 
	if limit == 0  || cursor+limit >= uint64(l){
		return allEntities[cursor:], nil
	}

	return allEntities[cursor:cursor+limit], nil
}

func (ts *TrackService) Create(title string) (uint64, error) {
	ts.LastID++
	allEntities = append(allEntities, logistic.NewTrack(ts.LastID, title))
	return ts.LastID, nil
}

func (ts *TrackService) Update(trackID uint64, track logistic.Track) error {
	return nil
}

func (ts *TrackService) Remove(trackID uint64) (bool, error) {
	var idx int
	var found bool

	for i, item := range(allEntities) {
		if item.ID == trackID {
			idx = i
			found = true
			break
		}
	}

	if !found {
		return false, fmt.Errorf("item not found")
	}

	allEntities = append(allEntities[:idx], allEntities[idx+1:]...)
	
	return true, nil
}
