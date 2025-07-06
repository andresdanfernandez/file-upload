#!/bin/bash

# Test script for Railway deployment
# Usage: ./scripts/test-deployment.sh [backend-url] [frontend-url]

BACKEND_URL=${1:-"http://localhost:8080"}
FRONTEND_URL=${2:-"http://localhost:5173"}

echo "🧪 Testing Railway Deployment"
echo "=============================="
echo "Backend URL: $BACKEND_URL"
echo "Frontend URL: $FRONTEND_URL"
echo ""

# Test backend health endpoint
echo "🔍 Testing Backend Health..."
HEALTH_RESPONSE=$(curl -s "$BACKEND_URL/health")
if [[ $HEALTH_RESPONSE == *"healthy"* ]]; then
    echo "✅ Backend is healthy"
else
    echo "❌ Backend health check failed"
    echo "Response: $HEALTH_RESPONSE"
fi

echo ""

# Test frontend accessibility
echo "🔍 Testing Frontend Accessibility..."
FRONTEND_RESPONSE=$(curl -s -I "$FRONTEND_URL" | head -n 1)
if [[ $FRONTEND_RESPONSE == *"200"* ]]; then
    echo "✅ Frontend is accessible"
elif [[ $FRONTEND_RESPONSE == *"404"* ]]; then
    echo "❌ Frontend returned 404 - check deployment"
else
    echo "⚠️  Frontend response: $FRONTEND_RESPONSE"
fi

echo ""

# Test CORS preflight
echo "🔍 Testing CORS Configuration..."
CORS_RESPONSE=$(curl -s -X OPTIONS -H "Origin: $FRONTEND_URL" \
    -H "Access-Control-Request-Method: POST" \
    -H "Access-Control-Request-Headers: Authorization" \
    "$BACKEND_URL/upload")

if [[ $CORS_RESPONSE == *"200"* ]] || [[ $CORS_RESPONSE == "" ]]; then
    echo "✅ CORS appears to be configured correctly"
else
    echo "⚠️  CORS might have issues"
fi

echo ""
echo "🎯 Next Steps:"
echo "1. Visit $FRONTEND_URL in your browser"
echo "2. Try to sign up/sign in"
echo "3. Test file upload functionality"
echo "4. Check browser console for any errors" 