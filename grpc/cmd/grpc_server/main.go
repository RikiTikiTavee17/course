package main

import (
	"context"
	"errors"
	"fmt"
	desc "github.com/RikiTIkiTavee17/course/grpc/pkg/note_v1"
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

var notes = &SyncMap{elems: make(map[int64]*desc.Note), num: 0}

type server struct {
	desc.UnimplementedNoteV1Server
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Note id: %d", req.GetId())
	n := notes.elems[req.GetId()]
	if n == nil {
		return &desc.GetResponse{
			Note: n,
		}, errors.New("note with such id not found")
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
		if req.Info.GetIsPublic() != nil {
			n.Info.IsPublic = req.Info.GetIsPublic().GetValue()
		}
		if req.Info.GetAuthor() != nil {
			n.Info.Author = req.Info.GetAuthor().GetValue()
		}
		if req.Info.GetTitle() != nil {
			n.Info.Title = req.Info.GetTitle().GetValue()
		}
		if req.Info.GetContent() != nil {
			n.Info.Content = req.Info.GetContent().GetValue()
		}
		n.UpdatedAt = timestamppb.New(time.Now())
	}
	return nil, nil
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
