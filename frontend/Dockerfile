# Этап сборки
FROM node:18-alpine AS builder

WORKDIR /app

# 1. Устанавливаем системные зависимости для сборки
RUN apk add --no-cache python3 make g++

# 2. Копируем только файлы зависимостей (для лучшего кэширования)
COPY package.json package-lock.json ./

# 3. Устанавливаем зависимости
RUN npm install --include=dev

# 4. Копируем исходный код
COPY . .

# 5. Варианты сборки с обработкой ошибок
RUN { \
  npm run build || \
  npx webpack --mode production || \
  ./node_modules/.bin/webpack --mode production || \
  (echo "Build failed" && exit 1); \
}

# Финальный образ
FROM node:18-alpine

WORKDIR /app

# 6. Копируем только необходимые файлы
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules

# 7. Устанавливаем serve (если не скопировался из builder)
RUN npm install serve

# 8. Запускаем приложение
CMD ["npx", "serve", "dist", "-l", "3000"]

EXPOSE 3000
HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD curl -f http://localhost:3000 || exit 1
