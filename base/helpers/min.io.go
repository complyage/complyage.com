//||------------------------------------------------------------------------------------------------||
//|| helpers/minio.go
//|| MinIO Upload, Download, Delete helpers
//||------------------------------------------------------------------------------------------------||

package helpers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

//||------------------------------------------------------------------------------------------------||
//|| Create a MinIO client (local helper)
//||------------------------------------------------------------------------------------------------||

func newMinioClient() (*minio.Client, error) {
	endpoint := os.Getenv("MINIO_ENDPOINT") // e.g. "localhost:9000"
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := os.Getenv("MINIO_USE_SSL") == "1"
	return minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
}

//||------------------------------------------------------------------------------------------------||
//|| Uploads a file (returns public URL)
//||------------------------------------------------------------------------------------------------||

func MinioUpload(bucket, objectName, contentType string, data []byte) (string, error) {
	client, err := newMinioClient()
	if err != nil {
		return "", err
	}
	ctx := context.Background()

	// Ensure bucket exists (ignore error if exists)
	err = client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucket := client.BucketExists(ctx, bucket)
		if errBucket != nil || !exists {
			return "", fmt.Errorf("bucket error: %w", err)
		}
	}
	_, err = client.PutObject(
		ctx,
		bucket,
		objectName,
		bytes.NewReader(data),
		int64(len(data)),
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return "", err
	}
	endpoint := os.Getenv("MINIO_PUBLIC_URL")
	url := fmt.Sprintf("%s/%s/%s", endpoint, bucket, objectName)
	return url, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Downloads an object (returns bytes, content-type, error)
//||------------------------------------------------------------------------------------------------||

func MinioDownload(bucket, objectName string) ([]byte, string, error) {
	client, err := newMinioClient()
	if err != nil {
		return nil, "", err
	}
	ctx := context.Background()
	obj, err := client.GetObject(ctx, bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", err
	}
	defer obj.Close()
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, obj); err != nil {
		return nil, "", err
	}
	stat, err := obj.Stat()
	if err != nil {
		return buf.Bytes(), "", nil // no content-type available
	}
	return buf.Bytes(), stat.ContentType, nil
}

//||------------------------------------------------------------------------------------------------||
//|| MinioDelete : Delete an object from MinIO
//||------------------------------------------------------------------------------------------------||

func MinioDelete(bucket, objectName string) error {
	client, err := newMinioClient()
	if err != nil {
		return err
	}
	ctx := context.Background()
	return client.RemoveObject(ctx, bucket, objectName, minio.RemoveObjectOptions{})
}
