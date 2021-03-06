package main

import (
	"flag"
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/giolekva/pcloud/pfs/api"
	"github.com/giolekva/pcloud/pfs/client"
)

var controllerAddress = flag.String("controller", "localhost:123", "Metadata storage address.")
var fileToUpload = flag.String("file", "", "File path to upload.")

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*controllerAddress, opts...)
	if err != nil {
		log.Fatalf("Failed to dial %s: %v", *controllerAddress, err)
	}
	defer conn.Close()
	uploader := client.NewFileUploader(api.NewMetadataStorageClient(conn))

	f, err := os.Open(*fileToUpload)
	if err != nil {
		panic(err)
	}

	uploader.Upload(f)
}
