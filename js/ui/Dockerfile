FROM node:14-alpine as frontend
WORKDIR /usr/src/app
COPY package.json yarn.lock ./
RUN yarn
COPY . .
RUN yarn build


FROM nginx:stable-alpine
COPY --from=frontend /usr/src/app/dist /usr/share/nginx/html
COPY nginx-frontend.conf.template /etc/nginx/templates/
RUN rm /etc/nginx/conf.d/default.conf