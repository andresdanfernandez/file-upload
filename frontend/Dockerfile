# Use Node.js 20 Alpine as base image
FROM node:20-alpine

# Set working directory
WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Install serve to run the built application
RUN npm install -g serve

# Expose port (will be overridden by Railway's PORT env var)
EXPOSE 4173

# Start production server
CMD ["serve", "-s", "dist", "-l", "4173"] 