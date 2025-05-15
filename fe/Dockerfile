
FROM node:24-alpine as build
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .

ARG VITE_API
ENV VITE_API=$VITE_API

RUN npm run build
EXPOSE 3000

CMD ["npx", "serve", "-s", "dist", "-l", "3000"]

