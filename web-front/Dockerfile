# Dockerfile for SvelteKit 
FROM node:21-alpine
WORKDIR /app
COPY package*.json ./
# TODO -- seperate dev and prod dependencies
RUN npm install
COPY . .
RUN npm run build

# # the sveltekit node adapter automatically runs on port 3000
# # to change the defaults, use:  HOST=127.0.0.1 PORT=4000 node build
# # https://kit.svelte.dev/docs/adapter-node
EXPOSE 3000
CMD ["node", "build"]
