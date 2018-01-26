package main

import (
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"fmt"
	"syscall"
	"google.golang.org/api/option"
	"io/ioutil"
	"github.com/simonjohansson/git-resource/cmd/common"
)

func main() {
	options := common.ParseOptions()

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(options.CredentialsPath))
	if err != nil {
		fmt.Println(err)
		syscall.Exit(-1)
	}

	bucket := client.Bucket(options.Bucket)
	obj := bucket.Object(options.Object)
	wr := obj.NewWriter(ctx)
	defer wr.Close()

	b, err := ioutil.ReadFile(options.FilePath)
	if err != nil {
		fmt.Print(err)
		syscall.Exit(-1)
	}

	_, err = wr.Write(b)
	if err != nil {
		fmt.Println(err)
		syscall.Exit(-1)
	}
	wr.Close()

	attr, err := obj.Attrs(ctx)
	if err != nil {
		fmt.Print(err)
		syscall.Exit(-1)
	}
	fmt.Println(attr.Generation)
}
