From alpine
RUN mkdir /app
COPY entrypoint.sh /app
WORKDIR /app
RUN chmod +x entrypoint.sh
COPY ./dist /app
EXPOSE 10011
CMD ["./entrypoint.sh","-"]