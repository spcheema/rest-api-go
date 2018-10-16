# rest-api-go

# Setup
- Dependencies are managed by usng `dep`. Please visit [https://github.com/golang/dep](https://github.com/golang/dep) for more information.
- Install gin build tool `go get github.com/codegangsta/gin`
- Create `.env` file and copy content from .env.dist file
- Update environment variables


## Run application
- run `dep ensure` to install required dependecies
- run `gin -p=8000 -a=8001 run main.go` and visit [localhost:8000](localhost:8000) in your browser to the application
> Run application without gin
- run `go run main.go` and visit [localhost:8001](localhost:8001) in your browser to the application.


## Todo
* [ ] write tests

## License
MIT © [Surinder Pal Singh](https://github.com/spcheema)