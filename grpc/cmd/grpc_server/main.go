package main

import (
	"context"
	"errors"
	"fmt"
	desc "github.com/RikiTikiTavee17/course/grpc/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"time"
)

const grpcPort = 50051

type SyncMap struct {
	elems map[int64]*desc.Note
	num   int64
}

type LoginMap struct {
	checks  map[string]string
	idLogin map[int64]string
	loginId map[string]int64
	num     int64
}

var notes = &SyncMap{elems: make(map[int64]*desc.Note), num: 1}
var users = &LoginMap{checks: make(map[string]string), idLogin: make(map[int64]string), loginId: make(map[string]int64), num: 1}

type server struct {
	desc.UnimplementedNoteV1Server
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Note id: %d", req.GetId())
	n := notes.elems[req.GetId()]
	if n == nil {
		return nil, errors.New("note with such id not found")
	}
	return &desc.GetResponse{
		Note: n,
	}, nil
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	var id int64 = notes.num
	notes.num++
	notes.elems[id] = &desc.Note{
		Id:        id,
		Info:      req.Info,
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
	}
	return &desc.CreateResponse{Id: id}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	if notes.elems[req.GetId()] == nil {
		return nil, errors.New("note with such id not found")
	} else {
		n := notes.elems[req.GetId()]
		if req.Info.GetAuthor() != nil {
			n.Info.Author = req.Info.GetAuthor().GetValue()
		}
		if req.Info.GetTitle() != nil {
			n.Info.Title = req.Info.GetTitle().GetValue()
		}
		if req.Info.GetContent() != nil {
			n.Info.Content = req.Info.GetContent().GetValue()
		}
		if req.Info.GetStatus() != nil {
			n.Info.Status = req.Info.GetStatus().GetValue()
		}
		if req.Info.GetDeadLine() != nil {
			n.Info.DeadLine = req.Info.GetDeadLine()
		}
		n.UpdatedAt = timestamppb.New(time.Now())
	}
	return nil, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if notes.elems[req.GetId()] == nil {
		return nil, errors.New("note with such id not found")
	} else {
		notes.elems[req.GetId()] = nil
	}
	return nil, nil
}

func (s *server) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	curr := make([]*desc.Note, 0)
	personId := req.GetPersonId()
	for id := range notes.elems {
		if notes.elems[id] != nil && personId == notes.elems[id].GetInfo().GetAuthor() && !notes.elems[id].GetInfo().GetStatus() {
			n := notes.elems[id]
			curr = append(curr, n)
		}
	}
	return &desc.ListResponse{Notes: curr}, nil
}

func (s *server) CreatePerson(ctx context.Context, req *desc.CreatePersonReqest) (*desc.CreatePersonResponse, error) {
	if users.checks[req.GetLogin()] == "" {
		id := users.num
		users.num++
		users.idLogin[id] = req.GetLogin()
		users.checks[req.GetLogin()] = req.GetPassword()
		users.loginId[req.GetLogin()] = id
		return &desc.CreatePersonResponse{Id: id}, nil
	} else {
		return nil, errors.New("user with this name is  already registered")
	}
}

func (s *server) LogInPerson(ctx context.Context, req *desc.LogInPersonRequest) (*desc.LogInPersonResponce, error) {
	if users.checks[req.GetLogin()] == req.Password {
		return &desc.LogInPersonResponce{Id: users.loginId[req.GetLogin()]}, nil
	} else {
		return nil, errors.New("incorrect login or password")
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
