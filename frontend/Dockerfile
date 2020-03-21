FROM node:13-slim

COPY . /src

RUN mkdir -p /app/node_modules

RUN cp -r /src/* /app/

WORKDIR /app

RUN npm install -g serve

VOLUME /app

EXPOSE 5000

CMD serve -s build
