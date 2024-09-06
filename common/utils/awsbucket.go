package utils

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadS3(title string,buf []byte)(*manager.UploadOutput, error){
	fileName:= fmt.Sprintf("%s.png",title)

	cfg, err:= config.LoadDefaultConfig(context.TODO())
	if err != nil{
		fmt.Println("wrong Config",err.Error())
		return nil,err
	}
fmt.Println("error", cfg)
	client := s3.NewFromConfig(cfg)

bucketName := "akkprojectimage"

uploader := manager.NewUploader(client)

input := &s3.PutObjectInput{
	Bucket: aws.String(bucketName),
	Key: aws.String(fileName),
	Body: bytes.NewReader(buf),
}

result, err := uploader.Upload(context.TODO(), input)
if err != nil{
	fmt.Println("upload Failure ",err.Error())
	return nil, fmt.Errorf("fail to upload image")
}

return result, nil
}