FROM node:18-alpine

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install --production  # Устанавливаем только production-зависимости
COPY . .

RUN npx webpack --mode production

EXPOSE 3000
CMD ["npx", "serve", "dist", "-l", "3000"]