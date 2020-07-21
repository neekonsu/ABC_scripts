FROM golang:1.14

RUN mkdir -p /go/src/github.com/neekonsu
WORKDIR /go/src/github.com/neekonsu

COPY ./dependencies.conf ./dependencies.conf

RUN export DEBIAN_FRONTEND=noninteractive
ARG DEBIAN_FRONTEND=noninteractive

RUN apt update -y
RUN apt install -y -q $(cat dependencies.conf) --fix-missing
RUN rm ./dependencies.conf

RUN git clone https://github.com/neekonsu/ABC-Enhancer-Gene-Prediction
RUN git clone  https://github.com/neekonsu/alphabet
RUN git clone https://github.com/neekonsu/juicer

RUN go get github.com/kr/pretty

RUN curl -O https://repo.anaconda.com/miniconda/Miniconda3-latest-Linux-x86_64.sh
RUN sh Miniconda3-latest-Linux-x86_64.sh
RUN conda config --add channels defaults
RUN conda config --add channels bioconda
RUN conda config --add channels conda-forge
RUN conda install pyranges

RUN chmod +x ./alphabet/*.sh

EXPOSE 5000