# Build stage
FROM node:18-alpine AS builder

# Set working directory
WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies with optimizations
RUN npm ci --silent --no-audit --no-fund

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Production stage
FROM node:18-alpine

# Install a minimal server
RUN npm install -g serve

# Set working directory
WORKDIR /app

# Copy built files
COPY --from=builder /app/dist ./dist

# Expose port
EXPOSE 3000

# Start the server
CMD ["serve", "-s", "dist", "-l", "3000"] 