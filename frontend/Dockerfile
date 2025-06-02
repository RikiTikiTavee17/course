# Используем многоэтапную сборку для уменьшения размера финального образа
FROM node:18-alpine AS builder

WORKDIR /app

# 1. Копируем только файлы зависимостей
COPY package.json package-lock.json ./

# 2. Устанавливаем зависимости (включая devDependencies)
RUN npm install --include=dev

# 3. Копируем исходный код
COPY . .

# 4. Собираем приложение
RUN npm run build  # Или webpack --mode production, если нет скрипта build

# 5. Устанавливаем serve локально (не глобально)
RUN npm install serve

# Финальный образ
FROM node:18-alpine

WORKDIR /app

# 6. Копируем только необходимые файлы из стадии builder
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules/serve ./node_modules/serve
COPY --from=builder /app/package.json ./

# 7. Устанавливаем только production зависимости
RUN npm install --omit=dev

# 8. Используем serve из node_modules
CMD ["npx", "serve", "dist", "-l", "3000"]

# Метаинформация
EXPOSE 3000
HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD curl -f http://localhost:3000 || exit 1
