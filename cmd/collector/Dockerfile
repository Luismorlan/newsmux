FROM public.ecr.aws/lambda/provided:al2 as build

# This docker file should have build context in project root folder
# eg: docker build -t data_collector -f Dockerfile ../../

# install compiler
RUN yum install -y golang
RUN go env -w GOPROXY=direct

# cache dependencies
ADD ./go.mod ./go.sum ./
RUN go mod download

# build
ADD . .
RUN GOOS=linux GOARCH=amd64 go build -o /main cmd/collector/main.go

# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /main /main

ARG ENV_ARG
ENV NEWSMUX_ENV=${ENV_ARG}

# switch to prod when roll out to prod
# dev will send all msgs into stderr
ENTRYPOINT [ "/main", "-service=collector" ]
