# Go-SearchEngine
 Search engine using Golang for school project made by Johnny HU and Raphael TURMEL

## How to use it
### 1. Install Go
First of all, you need to install Go on your computer. You can download it [here](https://golang.org/dl/).

### 2. Clone the project
Then, you need to clone the project on your computer. You can do it by using the following command line:
```
git clone
```

### 3. Build the project
Now, you need to build the project. You can do it by using the following command line:
```
go build -o . ./...
```

### 4. Run the project
Finally, you can run the project. You can do it by using the following command lines:
```
./index
```
Then
```
./search
```
And
```
./crawler
```

### 5. Use the project
To add a site :
```
curl -X POST http://localhost:8080/sites   -H 'Content-Type: application/json'   -d "{ ... }"
```
Example :
```
curl -X POST http://localhost:8080/sites   -H 'Content-Type: application/json'   -d "{\"ID\": 3, \"Hostip\": \"http://5.135.178.104:10987/\", \"Domain\": \"http://5.135.178.104:10987/\" }"
```

To get all sites :
```
curl -X GET http://localhost:8080/sites  -H 'Content-Type: application/json'
```

To get all files   :
```
curl -X GET http://localhost:8080/files  -H 'Content-Type: application/json'
```