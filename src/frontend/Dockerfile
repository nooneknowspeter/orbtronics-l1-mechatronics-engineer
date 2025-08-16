FROM oven/bun

WORKDIR /app

COPY package.json package.json
RUN bun install

COPY . .

COPY .env.development .env.development
COPY .env.production .env.production

RUN bun run build

CMD [ "bun", "run", "preview", "--host", "--mode", "production" ]