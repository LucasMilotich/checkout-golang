FROM "golang:rc-alpine3.12"
# 
RUN mkdir /app
ADD . /app/

WORKDIR /app
RUN  go build -v .
RUN adduser -S -D -H -h /app appuser
USER appuser

EXPOSE 8080
CMD ["./checkout-golang" ]
