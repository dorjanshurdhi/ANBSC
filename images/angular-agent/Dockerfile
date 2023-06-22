# Usa un'immagine di base con Node.js installato
FROM node:14

# Installa Git
RUN apt-get update && apt-get install -y git

# Imposta la directory di lavoro
WORKDIR /app

# Copia il file package.json e il file package-lock.json nella directory di lavoro
COPY package.json .
COPY package-lock.json .

# Installa le dipendenze
RUN npm install

# Aggiungi Angular CLI globalmente
RUN npm install -g @angular/cli

# Copia tutto il contenuto dell'app nella directory di lavoro
COPY . .

# Avvia l'app Angular
CMD ["ng", "serve", "--host", "0.0.0.0"]

# Espone la porta 4200 (la porta predefinita per l'app Angular)
EXPOSE 4200
