package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	domain "newnok6/logic-test/pie-fire-dire/internal/core/domain"
	"newnok6/logic-test/pie-fire-dire/internal/core/services"
	pb "newnok6/logic-test/pie-fire-dire/internal/proto"

	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	address    string
	listener   net.Listener
}

func NewGRPC(addr string) *Server {

	return &Server{
		grpcServer: grpc.NewServer(),
		address:    addr,
	}
}

func (s *Server) Start() {
	var err error
	s.listener, err = net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("gRPC listen error: %v", err)
	}

	pb.RegisterMeatServiceServer(s.grpcServer, &service{}) // Implement your service
	log.Printf("gRPC server running at %s\n", s.address)

	if err := s.grpcServer.Serve(s.listener); err != nil {
		log.Fatalf("gRPC serve error: %v", err)
	}
}

func (s *Server) Stop() {
	log.Println("Stopping gRPC server...")
	s.grpcServer.GracefulStop()
}

type service struct {
	pb.UnimplementedMeatServiceServer
}

func (s *service) GetMeat(ctx context.Context, req *pb.MeatRequest) (*pb.MeatReply, error) {
	// Implement your logic here
	// For example, you can return a dummy response
	fileName := "file.txt"
	filePath := "/Users/panupak/Projects/logic-test/pie-fire-dire/files"
	fileType := "txt"

	fileMeta := domain.FileMeta{
		FileName: fileName,
		FilePath: filePath,
		FileType: fileType,
	}

	processFileService := services.NewProcessFileService(fileMeta)

	fmt.Println("File Name:", processFileService.GetFileName())
	fmt.Println("File Path:", processFileService.GetFilePath())

	meatList, err := processFileService.GetMeatList()

	if err != nil {
		log.Printf("Error getting meat list: %v", err)
		return nil, err
	}
	reply := &pb.MeatReply{
		Beef: meatList,
	}
	return reply, nil
}

// GetMeat implements the GetMeat method required by the MeatServiceServer interface.
