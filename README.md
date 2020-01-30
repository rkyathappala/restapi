# REST API
REST API for file management, intended to support creating files and driectories, reading files, writing files, and listing contents f directories .

## Create file
Whiile using API testing client such as Postman, make a POST request with URL= http://localhost:8000/api/makeFile. The post should expect a JSON entry of the form:

{
  "_id": "500100",
  "dest_id": "1000",
  "name": "file1.txt",
	"content": "This is a file continaing my thoughts."
}

## Create directory
Make a POST request with URL= http://localhost:8000/api/makeFolder. The post should expect a JSON entry of the form:

{
  "_id": "1050",
  "dest_id": "1000",
	"name": "dumbDir.txt",
	"subs": [],
	"files": []
}

