FROM golang:1.9.2-stretch
MAINtAINER yskn67

RUN go get gonum.org/v1/gonum/blas
RUN go get -d gonum.org/v1/netlib; exit 0
RUN cd / \
    && git clone https://github.com/xianyi/OpenBLAS \
    && cd OpenBLAS \
    && make \
    && make install \
    && echo "/opt/OpenBLAS/lib" >> /etc/ld.so.conf.d/libc.conf \
    && ldconfig \
    && cd /go
RUN CGO_LDFLAGS="-L/opt/OpenBLAS/lib -lopenblas" go install gonum.org/v1/netlib/blas/netlib
RUN git clone https://github.com/yskn67/gonum_compare.git
RUN cd /go/gonum_compare && go test -bench .
