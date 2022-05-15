FROM node:alpine

WORKDIR /app

COPY . .

RUN npm i && chmod 777 -R ./node_modules

EXPOSE 8080

CMD ["npm", "run", "serve"]