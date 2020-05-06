
FROM scratch

RUN mkdir -p /usr/local/rfid/xlslr5603
COPY .env /usr/local/rfid/xlslr5603
COPY ./server /usr/local/rfid/xlslr5603/server

RUN chmod 777 /usr/local/rfid/xlslr5603/server

WORKDIR /usr/local/rfid/xlslr5603

EXPOSE 80

CMD ["/usr/local/rfid/xlslr5603/server"]