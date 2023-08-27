FROM centos:latest
MAINTAINER "Tom Gur"
WORKDIR /etc/yum.repos.d
RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*
RUN ls -la /etc/yum.repos.d/
RUN yum install -y epel-release
RUN yum install -y golang
COPY main.go /go/src/
EXPOSE 8080
ENTRYPOINT ["go", "run", "/go/src/main.go"]