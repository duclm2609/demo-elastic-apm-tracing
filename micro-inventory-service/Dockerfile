FROM openjdk:11.0.6-jre-slim-buster
ENV TZ=Asia/Ho_Chi_Minh
WORKDIR /app

COPY apm/elastic-apm-agent-1.13.0.jar /app/elastic-apm-agent.jar
COPY apm/elasticapm.properties /app/elasticapm.properties

COPY ./target/*.jar /app/micro-inventory.jar
CMD exec java -javaagent:elastic-apm-agent.jar $JAVA_OPTS -jar micro-inventory.jar