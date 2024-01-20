# Use an official Nginx image as a parent image
FROM nginx:stable-alpine

# Set the working directory in the container to /usr/share/nginx/html
WORKDIR /usr/share/nginx/html

# Remove the default Nginx index.html
RUN rm -rf ./*

# Copy the static content from web-front/build directory to the working directory in the container
COPY build .

# Change the ownership of all files in /usr/share/nginx/html to nginx
RUN chown -R nginx:nginx .

# Expose port 80
EXPOSE 80

# Start Nginx when the container launches
CMD ["nginx", "-g", "daemon off;"]