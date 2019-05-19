# http-post

Go version HTTP server for handle POST request, retrieve the data as string and send it into a buffered channel. The POST request is mocked by a Python script, data comes from sample json file.

## Usage

```bash
# Console 1: launch HTTP server in Go
go run http_server.go
# Console 2: mock POST request in Python, load data from sample json file
python mock_post_request.py
# Back to Console 1 of Go HTTP server, it should print the string format of data
```

