# Usa un'immagine base con JDK 17
FROM eclipse-temurin:17-jdk

# Installa Maven
RUN apt-get update && apt-get install -y maven git

# Imposta la variabile d'ambiente per Maven
ENV MAVEN_HOME /usr/share/maven

# Aggiungi Maven al PATH
ENV PATH $MAVEN_HOME/bin:$PATH
