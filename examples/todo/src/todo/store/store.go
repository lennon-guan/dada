package store

import "time"

type TodoRec struct {
	Id       int
	Title    string
	Content  string
	Done     bool
	CreateAt time.Time
	FinishAt time.Time
}

type todoStore struct {
	records []*TodoRec
	lastId  int
}

func NewStore() *todoStore {
	return &todoStore{
		records: []*TodoRec{},
		lastId:  0,
	}
}

func (s *todoStore) Add(rec *TodoRec) {
	s.lastId++
	rec.Id = s.lastId
	s.records = append(s.records, rec)
}

func (s *todoStore) List() []*TodoRec {
	return s.records
}

func (s *todoStore) FindById(id int) (index int, rec *TodoRec) {
	for index, rec = range s.records {
		if rec.Id == id {
			return
		}
	}
	return -1, nil
}

var TodoStore *todoStore

func init() {
	TodoStore = NewStore()
}
