# s3-uploader
This project can be used to upload file or folder (can have nested folder also) to AWS s3 storage.
This takes very minimum time to upload large folders. 
approx 450mb takes 2m

## Environment variables
Change aws configs in app.env file

## Run command
go run main.go <source_path> <foldername>

## Create exe and execute
go build
go install
s3-uploader <source_path> <foldername>
s3-uploader will be the exe name

