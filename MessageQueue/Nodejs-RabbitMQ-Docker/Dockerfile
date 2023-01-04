FROM node:18-alpine as base
WORKDIR /app
COPY . .
RUN yarn


FROM base as development
RUN apk add --no-cache bash
RUN wget -O /bin/wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
RUN chmod +x /bin/wait-for-it.sh


FROM base as production
ENV NODE_ENV=production
RUN yarn build
COPY ./dist ./
CMD ["node", "dist/consumer.js"]
