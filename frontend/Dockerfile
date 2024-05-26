FROM node:20-alpine

WORKDIR /app

# pre req is to build this on local and then copy
# dont want bloated images due to npm install
# mid way to do things
COPY build /app/build

RUN npm install -g http-server

CMD ["http-server", "build", "-p", "3000"]

EXPOSE 3000