FROM "golang:rc-alpine3.12"

COPY checkout-golang .

EXPOSE 8080

CMD [ "./checkout-golang" ]
