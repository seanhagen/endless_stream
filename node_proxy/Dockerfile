FROM node:8-alpine

RUN npm install -g nodemon

# Create app directory
RUN mkdir -p /app
WORKDIR /app

# Install dependencies
COPY package.json .
COPY package-lock.json .

RUN npm install grpc
RUN npm install
RUN npm rebuild

# Bundle app source
COPY index.js .
COPY proto .

# Exports
EXPOSE 3000
CMD [ "nodemon" ]
