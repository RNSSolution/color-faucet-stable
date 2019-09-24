FROM node:10-alpine

RUN apk add git

WORKDIR /go/src/project/

COPY . /go/src/project/

RUN npm install -g yarn

RUN yarn version

COPY package.json ./

RUN yarn install

CMD ["yarn","serve:dist"]
