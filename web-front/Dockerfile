# Use an official Node image as a parent image
FROM node:latest as builder

# Set the working directory in the container to /app
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the source code to the working directory
COPY . .

# Build the project
RUN npm run build

# Use an official Nginx image as a parent image
FROM nginx:stable-alpine

# Set the working directory in the container to /usr/share/nginx/html
WORKDIR /usr/share/nginx/html

# Remove the default Nginx index.html
RUN rm -rf ./*

# Copy the static content from the builder stage to the working directory in the container
COPY --from=builder /app/build .

# Change the ownership of all files in /usr/share/nginx/html to nginx
RUN chown -R nginx:nginx .

# Expose port 80
EXPOSE 80

# Start Nginx when the container launches
CMD ["nginx", "-g", "daemon off;"]