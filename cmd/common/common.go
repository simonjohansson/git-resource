package common

import (
	"flag"
	"cloud.google.com/go/storage"
)

type Options struct {
	Bucket          string
	Object          string
	Generation      int64
	CredentialsPath string
	FilePath        string
}

func ParseOptions() Options {
	bucket := flag.String("bucket", "", "Bucket to look for a file")
	object := flag.String("object", "", "Object in the bucket")
	generation := flag.Int64("generation", 0, "Generation of the file")
	credentialsPath := flag.String("credentialsPath", "", "Path to file holding credentials")
	filePath := flag.String("filePath", "", "Path to file to read or write")

	flag.Parse()

	return Options{
		Bucket:          *bucket,
		Object:          *object,
		Generation:      *generation,
		CredentialsPath: *credentialsPath,
		FilePath:        *filePath,
	}
}

func (o Options) ConstructObject(client *storage.Client) storage.ObjectHandle {
	bucket := client.Bucket(o.Bucket)

	if o.Generation != 0 {
		return *bucket.Object(o.Object).Generation(o.Generation)
	}
	return *bucket.Object(o.Object)
}
