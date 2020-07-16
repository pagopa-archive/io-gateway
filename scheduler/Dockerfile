FROM node:10

ENV PORT 3100
ENV  WSK_INSTALL https://github.com/apache/openwhisk-cli/releases/download/1.0.0/OpenWhisk_CLI-1.0.0-linux-amd64.tgz
ENV BIN /usr/local/bin
EXPOSE 3100

COPY package.json package.json
RUN npm install

COPY . .
RUN npm run build

RUN curl -sL $WSK_INSTALL > $BIN/wsk.tgz
RUN tar xzvf $BIN/wsk.tgz -C $BIN  wsk

CMD ["node", "dist/"]
