FROM debian:bookworm

WORKDIR /work

# Install curl (required for sdkman)
RUN apt-get update && apt-get install -y curl zip unzip openssh-client

# Install sdkman
RUN curl -s "https://get.sdkman.io" | bash

# Set up environment for sdkman
ENV SDKMAN_DIR="/root/.sdkman"
ENV PATH="${PATH}:${SDKMAN_DIR}/bin"
SHELL ["/bin/bash", "-c"]

# Install Java 21 and Maven using sdkman
RUN set -e; \
    source "${SDKMAN_DIR}/bin/sdkman-init.sh" && \
    sdk install java 21.0.1-open && \
    sdk install maven

# Set JAVA_HOME
ENV JAVA_HOME="${SDKMAN_DIR}/candidates/java/current"
ENV PATH="${JAVA_HOME}/bin:${PATH}"

COPY jmh_example_benchmark .
RUN ${SDKMAN_DIR}/candidates/maven/current/bin/mvn clean package

COPY boot.sh .

RUN rm -rf /root/.ssh
RUN ln -s /shared/ssh /root/.ssh

# CMD ["java", "-jar", "target/benchmarks.jar"]
CMD ["./boot.sh"]
