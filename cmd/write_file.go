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
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("/tmp/creds"))
	if err != nil {
		fmt.Println(err)
		syscall.Exit(-1)
	}

	obj := options.ConstructObject(client)
	wr := obj.NewWriter(ctx)
	defer wr.Close()

	b, err := ioutil.ReadFile(options.FilePath) // just pass the file name
	if err != nil {
		fmt.Print(err)
		syscall.Exit(-1)
	}

	_, err = wr.Write(b)
	if err != nil {
		fmt.Println(err)
		syscall.Exit(-1)
	}

	attr, _ := obj.Attrs(ctx)
	fmt.Println(attr.Generation)
}
