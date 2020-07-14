FROM node:10

ENV PORT 3100
EXPOSE 3100

COPY package.json package.json
RUN npm install

COPY . .
RUN npm run build

CMD ["node", "dist/"]
