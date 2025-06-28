FROM node:20

WORKDIR /app

# Сначала копируем только package.json и package-lock.json (если он есть)
COPY package*.json ./

# Устанавливаем зависимости
RUN npm install && npm i

# Копируем остальные файлы
COPY . .

# Запускаем приложение
CMD ["npm", "run", "dev"]
