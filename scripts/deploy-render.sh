#!/bin/bash

# Render Deployment Helper Script
# This script helps you test your application locally before deploying to Render

echo "🚀 Render Deployment Helper"
echo "=========================="

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker is not running. Please start Docker and try again."
    exit 1
fi

echo "✅ Docker is running"

# Build and test backend
echo ""
echo "🔨 Building Backend..."
cd "$(dirname "$0")/.."
docker build -t file-upload-backend .

if [ $? -eq 0 ]; then
    echo "✅ Backend build successful"
else
    echo "❌ Backend build failed"
    exit 1
fi

# Build and test frontend
echo ""
echo "🔨 Building Frontend..."
cd frontend
docker build -t file-upload-frontend .

if [ $? -eq 0 ]; then
    echo "✅ Frontend build successful"
else
    echo "❌ Frontend build failed"
    exit 1
fi

echo ""
echo "🎉 All builds successful!"
echo ""
echo "📋 Next Steps:"
echo "1. Push your code to GitHub"
echo "2. Follow the deployment guide in RENDER_DEPLOYMENT.md"
echo "3. Set up your environment variables in Render dashboard"
echo ""
echo "🔗 Useful Links:"
echo "- Render Dashboard: https://dashboard.render.com"
echo "- Deployment Guide: RENDER_DEPLOYMENT.md" 