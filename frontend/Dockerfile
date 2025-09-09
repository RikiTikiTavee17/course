# =========================
# Этап сборки
# =========================
FROM node:18-alpine AS builder

WORKDIR /app

RUN apk add --no-cache python3 make g++

# !!! Обязательно указываем NODE_ENV=development, чтобы ставились devDependencies
ENV NODE_ENV=development

COPY package*.json ./
RUN npm ci --include=dev

COPY . .

# Проверка, что webpack реально установлен
RUN ls -la node_modules/.bin && npm list --depth=0

# Запуск сборки напрямую через Node
RUN node ./node_modules/webpack/bin/webpack.js --mode production

# =========================
# Финальный образ
# =========================
FROM node:18-alpine

WORKDIR /app

COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules

RUN npm install serve

CMD ["npx", "serve", "dist", "-l", "3000"]

EXPOSE 3000
HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD curl -f http://localhost:3000 || exit 1
