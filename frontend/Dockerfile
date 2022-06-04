FROM node:17.3.0-alpine

WORKDIR /app

# install app dependencies
COPY package.json ./
COPY package-lock.json ./
RUN npm install 

# add app
COPY . ./

# start app
CMD ["npm", "start"]