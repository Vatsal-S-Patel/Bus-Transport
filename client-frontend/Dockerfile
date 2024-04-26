FROM node:lts-bookworm-slim AS build-image

WORKDIR /frontend

COPY /public ./public/
COPY /src ./src/
COPY /package.json .

EXPOSE 3000

RUN npm i

RUN npm run build

COPY nginx.conf ./build/

FROM nginx:alpine

COPY --from=build-image /frontend/build /usr/share/nginx/html

RUN rm /etc/nginx/conf.d/default.conf
RUN mv /usr/share/nginx/html/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]