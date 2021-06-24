package server

import (
	"net/http"
	"sort"
	"strings"
	"sync"
	"trendings/server/page"
	"trendings/trends"
)

type Server struct {
	data    sync.Map
	service trends.Trends24
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	http.HandleFunc("/", s.index)
	http.HandleFunc("/data", s.readData)
	return http.ListenAndServe(":8080", nil)
}

func (s *Server) index(w http.ResponseWriter, _ *http.Request) {
	w.Write(page.Content)
}

func (s *Server) readData(w http.ResponseWriter, r *http.Request) {
	resp := toResponse(&s.data)
	jsonData(w, http.StatusOK, resp)
}

func (s *Server) PublishResult(location string, hashtags []string) {
	s.data.Store(location, hashtags)
}

func (s *Server) CleanAll() {
	s.data.Range(func(key, value interface{}) bool {
		s.data.Delete(key)
		return true
	})
}

func toResponse(m *sync.Map) *DataResponse {
	type kv struct {
		Key   string
		Value []string
	}

	var ss []kv
	m.Range(func(key, value interface{}) bool {
		ss = append(ss, kv{key.(string), value.([]string)})
		return true
	})

	sort.Slice(ss, func(i, j int) bool {
		return strings.Compare(ss[i].Key, ss[j].Key) == -1
	})

	var resp DataResponse
	for _, v := range ss {
		resp.Trends = append(resp.Trends, &DataTrend{Location: v.Key, Hashtags: v.Value})
	}

	return &resp
}
