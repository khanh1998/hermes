FROM node:16-alpine as builder

ENV NODE_ENV build

WORKDIR /usr/src/app

COPY . .

RUN npm ci \
    && npm run build \
    && npm prune --production

# ---

FROM node:16-alpine

ENV NODE_ENV production

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/package*.json .
COPY --from=builder /usr/src/app/node_modules/ ./node_modules/
COPY --from=builder /usr/src/app/dist/ ./dist/

CMD ["node", "dist/server.js"]