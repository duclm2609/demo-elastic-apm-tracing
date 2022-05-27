FROM maven:3-adoptopenjdk-11 as builder
WORKDIR /src
COPY . .
RUN --mount=type=cache,target=/root/.m2 mvn -T 1C -U -B clean package -DskipTests

FROM openjdk:11.0.6-jre-slim-buster
ENV TZ=Asia/Ho_Chi_Minh
WORKDIR /app

ARG APM_VERSION
RUN apt update && apt install -y curl
RUN curl -o apm.jar https://search.maven.org/remotecontent?filepath=co/elastic/apm/elastic-apm-agent/$APM_VERSION/elastic-apm-agent-$APM_VERSION.jar

COPY --from=builder /src/target/*.jar /app/micro-inventory.jar
CMD exec java -javaagent:apm.jar $JAVA_OPTS -jar micro-inventory.jar