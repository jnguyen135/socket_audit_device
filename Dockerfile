FROM centos:7

MAINTAINER James Nguyen <jnguyen@peernova.com>

RUN rpm --import https://mirror.go-repo.io/centos/RPM-GPG-KEY-GO-REPO && \
    curl -s https://mirror.go-repo.io/centos/go-repo.repo | tee /etc/yum.repos.d/go-repo.repo && \
    yum -y update && \
    yum -y install golang && \
    yum clean all

COPY . /app

WORKDIR /app

RUN go build -o main .

ENTRYPOINT ["/app/main"]
