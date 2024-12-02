FROM node:16-alpine

WORKDIR /usr/src/app

COPY package*.json ./
COPY nodemon.json ./

RUN rm -rf node_modules
RUN npm install
RUN npm install -g nodemon

COPY . .

EXPOSE 8080

CMD [ "npm", "run", "dev" ]