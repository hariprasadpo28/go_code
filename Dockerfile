FROM node:16.4.2-alpine3.12
ENV NODE_ENV production
# RUN pwd
RUN mkdir /app && chown -R node:node /app
# #to copy manifest files from inventory repo to current working directory
# RUN mkdir /app/inventory && chown -R node:node /app/inventory
# RUN mkdir /app/inventory/build && chown -R node:node app/inventory/build
# RUN cd Security/ && ls -l
RUN pwd
# RUN cd /home/runner/work/ && ls -l
RUN ls -l
# RUN cd /home/runner/work && ls -l
RUN find / -type f -name getalerts.js
# COPY getalerts.js app/
# COPY calculator.js app/
WORKDIR /app
RUN ls -l
USER node
RUN node getalerts.js
# COPY *.js /app
# COPY inventory/build/*.props /app/inventory/build


EXPOSE 8080
CMD ["node", "getalerts.js"]