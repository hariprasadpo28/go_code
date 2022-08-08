FROM node:16.4.2-alpine3.12
ENV NODE_ENV production
# RUN pwd
RUN mkdir /app && chown -R node:node /app
RUN pwd
RUN ls -l
WORKDIR /app
RUN ls -l
COPY main_repo/testing/test.txt /app
COPY cloned_repo/build.cmd /app
COPY main_repo/script.js /app
COPY main_repo/calculator.js /app
USER node
RUN node script.js
# COPY *.js /app
# COPY inventory/build/*.props /app/inventory/build


EXPOSE 8080
CMD ["node", "script.js"]