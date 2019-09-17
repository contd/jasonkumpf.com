FROM kumpfjn/serve:latest

RUN mkdir -p /code

WORKDIR /code

VOLUME ["/code"]

EXPOSE 5000/tcp

CMD ["serve","-l","5000"]
