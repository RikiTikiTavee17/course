# Этап сборки
FROM node:18-alpine AS builder

WORKDIR /app

# 1. Устанавливаем системные зависимости для сборки
RUN apk add --no-cache python3 make g++

# 2. Копируем только файлы зависимостей (для лучшего кэширования)
COPY package.json package-lock.json ./

# 3. Устанавливаем зависимости (включая dev-зависимости)
RUN npm install --include=dev

# 4. Копируем исходный код
COPY . .

# 5. Выставляем права на webpack и запускаем сборку
RUN chmod +x ./node_modules/.bin/webpack && \
    npm run build

# Финальный образ
FROM node:18-alpine

WORKDIR /app

# 6. Копируем только необходимые файлы
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules

# 7. Устанавливаем serve (на всякий случай, если нужно обновить)
RUN npm install serve

# 8. Запускаем приложение
CMD ["npx", "serve", "dist", "-l", "3000"]

EXPOSE 3000
HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD curl -f http://localhost:3000 || exit 1
