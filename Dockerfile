FROM node:16.15-alpine3.14 AS build

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 4444

RUN npm i pm2 -g

ENTRYPOINT ["pm2-runtime", "index.js"]