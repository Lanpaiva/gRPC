package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/lanpaiva/grpc/internal/database"
	"github.com/lanpaiva/grpc/internal/pb"
	"github.com/lanpaiva/grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		fmt.Printf("err to open database, %v", err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("err to open server, %v", err)
	}

	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
