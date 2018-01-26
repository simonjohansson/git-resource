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
	rc, err := obj.NewReader(ctx)
	if err != nil {
		fmt.Println(fmt.Sprintf("readFile: unable to open file %q, with generation %i from bucket %q %v", options.Object, options.Generation, options.Bucket, err))
		syscall.Exit(-1)
	}
	defer rc.Close()

	content, err := ioutil.ReadAll(rc)
	if err != nil {
		fmt.Println(err)
		syscall.Exit(-1)
	}

	err = ioutil.WriteFile(options.FilePath, content, 0644)
	if err != nil {
		fmt.Println(err)
		syscall.Exit(-1)
	}
}
