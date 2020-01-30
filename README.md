# REST API
REST API for file management, intended to support creating files and driectories, reading files, writing files, and listing contents of directories .

## Starting the API
Open a golang environment, and run the following command.

go run main.go

## Create file
While using an API testing client such as Postman, make an HTTP POST request with URL= http://localhost:8000/api/makeFile. The post should expect a JSON entry of the form:

{	"_id": "500100",
	"dest_id": "1000",
	"name": "file1.txt",
	"content": "This is a file continaing my thoughts."
}

If you intend on setting the file's ID manually, the "_id" key should become "id".

## Create directory
Make an HTTP POST request with URL= http://localhost:8000/api/makeFolder. The post should expect a JSON entry of the form:

{
	"_id": "1050",
  	"dest_id": "1000",
	"name": "dumbDir.txt",
	"subs": [],
	"files": []
}
If you intend on setting the folder's ID manually, the "_id" key should become "id".

## List contents of directory
Listing contents of the home directory (in which all files and subdirectories are nested) is done by opening the api, which is done by making an HTTP GET request from URL= http://localhost:8080/api

Listing contents of a specific directory is one by making an HTTP GET request from URL= http://localhost:8080/api/getFolder/{ID}. In order to address a specific file or directory to view, you must know that item's ID. These IDs can be seen in the home directory, which is explained just above.

## Read files
Make an HTTP GET request with URL= http://localhost:8000/api/getFile/{ID}. 

## Writing files
To write to (or update) a file, make an HTTP PUT request to URL= http://localhost:8000/updateFile/{ID}. Similar to the procedure for making a file, use the JSON template for the PUT request. 

The write method works by first deleting the existing file data, then replacing the all fields except the ID field, which is the only unique, immutable element. This method also allows for the file to simply be moved by changing the file's destination ID element.

**_The delete method is not currently working, thus impacting the updae method_**
