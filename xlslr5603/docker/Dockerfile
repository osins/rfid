
FROM scratch

RUN mkdir -p /usr/local/rfid/xlslr5603
ADD ./server /usr/local/rfid/xlslr5603/

RUN chmod 777 /usr/local/rfid/xlslr5603/server

WORKDIR /usr/local/rfid/xlslr5603

EXPOSE 80

CMD ["/usr/local/rfid/xlslr5603/server"]