FROM golang



# create a working directory
WORKDIR /go/src/app


# add source code
ADD RestFullAndGorm RestFullAndGorm


EXPOSE 9191


# run main.go
CMD ["go", "run", "RestFullAndGorm/src/webservice/main.go"]
