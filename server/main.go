package main

import (
	"context"
	"log"
	"net"

	pb "github.com/RaminCH_self/Go3_gRPC/lec6/server/proto/consigment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50001"
)

//obyekt kotoriy budet xranit info o vnutrennix konfiguratsiyax
type repository interface {
	Create(*pb.Command) (*pb.Command, error)
	GetAll() []*pb.Command
}

//Repository... Nasha DB - localnaya - v dalneyshem realniye DB na Docker budut
type Repository struct {
	commands []*pb.Command
}

//Create...		(etot metod budet delat Create dla 'type service struct' )
func (r *Repository) Create(command *pb.Command) (*pb.Command, error) {
	updatedCommands := append(r.commands, command)
	r.commands = updatedCommands //yesli prisvoit srazu bez 'updatedCommands' to budet infinite loop !
	return command, nil
}

//GetAll...
func (r *Repository) GetAll() []*pb.Command {
	return r.commands
}

type service struct {
	repo repository //u servisa budet yedinstvennoye pole -> gde xranatsa danniye
}

// repo - eto nekiy obj., udovl. interfeysu repository, kotoriy umeyet delat Create
// takje obj. tipa 'service' doljen udovl interfeysu, kot nax. v consigment.proto -> service ShippingService { rpc CreateCommand(Command) ...
// type service struct -> doljen umet Create i v cons..pb.go -> naxodim (unimplemented) CreateCommand... i kopiruyem suda(chutok izmeniv) --> see below

func (s *service) CreateCommand(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	command, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	log.Printf("Request to create response: %v", command)
	return &pb.Response{Created: true, Command: command}, nil //Response -> check consigment.proto and consigment.pb.go
}

//GetAllCommands...
func (s *service) GetAllCommands(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	commands := s.repo.GetAll()
	return &pb.Response{Commands: commands}, nil
}

func main() {
	repo := &Repository{} //local storage

	//nastroyka gRPC servera
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen port: %v", err)
	}

	server := grpc.NewServer()

	//Registriruyem nash servis dla servera
	ourService := &service{repo}                         //repo v -> type service struct{...}
	pb.RegisterShippingServiceServer(server, ourService) //sopostavlayem (s *grpc.Server-grpc) s (srv ShippingServiceServer-nash) see con..pb.go file

	//chtobi vixodniye parametri servera soxranalis v go-runtime
	reflection.Register(server) //reflektim, chtobi danniye ne provalivalis v 'run time'

	log.Println("gRPC server runs on port: ", port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve from port: %v", port)
	}
}
