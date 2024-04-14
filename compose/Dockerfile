FROM mongo:6.0.14

ENV MONGO_INITDB_ROOT_USERNAME=mongo_root
ENV MONGO_INITDB_ROOT_PASSWORD=mongo_root
ENV APP_USER=app_user
ENV APP_PWD=app_user
ENV DB_NAME=test
ENV DB_COLLECTION_NAME=test
ENV MONGO_HOSTNAME=mongodb
ENV MONGO_PORT=27017

# Copy the initialization script into the container
COPY ./compose/mongo/init_db.js /docker-entrypoint-initdb.d/

CMD ["mongod", "--bind_ip_all"]