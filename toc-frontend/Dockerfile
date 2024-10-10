FROM node:lts-alpine

WORKDIR /app

COPY package.json ./package.json
COPY package-lock.json ./package-lock.json
COPY src ./src/
COPY public ./public/

RUN npm install
RUN npm run build
RUN tar -C build -czvf toc-frontend.tar.gz .