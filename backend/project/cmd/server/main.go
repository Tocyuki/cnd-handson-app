package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/internal/project/handler"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/internal/project/model"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/project/pkg/db"
	projectpb "github.com/cloudnativedaysjp/cnd-handson-app/backend/project/proto"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: project-service [server|migrate|reset]")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "server":
		runServer()
	case "migrate":
		runMigrate()
	case "reset":
		resetDB()
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}

func runServer() {
	// データベース接続初期化
	_, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// gRPCサーバーの設定
	port := os.Getenv("PORT")
	if port == "" {
		port = "50053" // プロジェクトサービス用のポート
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPCサーバーの作成
	grpcServer := grpc.NewServer()

	// gRPCサービスの登録
	projectService := &handler.ProjectServiceServer{}
	projectpb.RegisterProjectServiceServer(grpcServer, projectService)

	log.Printf("gRPC server listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func runMigrate() {
	conn, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.MigrateDB(conn, &model.Project{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}

func resetDB() {
	conn, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.ResetDB(conn, &model.Project{})
	if err != nil {
		log.Fatalf("Database reset failed: %v", err)
	}

	log.Println("Database reset completed successfully")
}
