# Render Deployment Guide

This guide will help you deploy your file upload application to Render.

## Prerequisites

1. **GitHub Repository**: Make sure your code is pushed to a GitHub repository
2. **Render Account**: Sign up at [render.com](https://render.com)
3. **AWS S3 Bucket**: You'll need your AWS credentials and S3 bucket name

## Step 1: Deploy Database

1. Go to your Render dashboard
2. Click "New +" → "PostgreSQL"
3. Configure:
   - **Name**: `file-upload-db`
   - **Database**: `fileupload`
   - **User**: `fileupload_user`
   - **Plan**: Free
4. Click "Create Database"
5. Note down the connection details (you'll need these for environment variables)

## Step 2: Deploy Backend

1. In Render dashboard, click "New +" → "Web Service"
2. Connect your GitHub repository
3. Configure the service:
   - **Name**: `file-upload-backend`
   - **Environment**: Docker
   - **Region**: Oregon (or closest to you)
   - **Branch**: `main`
   - **Root Directory**: Leave empty (root of repo)
   - **Dockerfile Path**: `Dockerfile`
   - **Health Check Path**: `/health`

### Environment Variables

Add these environment variables in the Render dashboard:

**Database Variables** (from your PostgreSQL service):
- `DB_HOST`: Your database host
- `DB_PORT`: Your database port (usually 5432)
- `DB_USER`: Your database user
- `DB_PASSWORD`: Your database password
- `DB_NAME`: `fileupload`

**AWS Variables**:
- `AWS_REGION`: Your AWS region (e.g., `us-east-1`)
- `AWS_ACCESS_KEY_ID`: Your AWS access key
- `AWS_SECRET_ACCESS_KEY`: Your AWS secret key
- `S3_BUCKET_NAME`: Your S3 bucket name

4. Click "Create Web Service"

## Step 3: Deploy Frontend

1. In Render dashboard, click "New +" → "Web Service"
2. Connect your GitHub repository
3. Configure the service:
   - **Name**: `file-upload-frontend`
   - **Environment**: Docker
   - **Region**: Oregon (or closest to you)
   - **Branch**: `main`
   - **Root Directory**: `frontend`
   - **Dockerfile Path**: `Dockerfile`

### Environment Variables

Add this environment variable:
- `VITE_API_BASE`: `https://your-backend-service-name.onrender.com`

4. Click "Create Web Service"

## Step 4: Update CORS (if needed)

If your frontend URL is different from the one in the code, update the CORS configuration in `cmd/main.go`:

```go
AllowOrigins: []string{
    "http://localhost:5173", 
    "http://localhost:3000", 
    "http://localhost:4173",
    "https://your-frontend-service-name.onrender.com",
},
```

## Step 5: Test Your Deployment

1. Wait for both services to deploy (green status)
2. Visit your frontend URL
3. Test file upload functionality
4. Check the backend health endpoint: `https://your-backend-service-name.onrender.com/health`

## Troubleshooting

### Common Issues:

1. **Database Connection**: Make sure all database environment variables are set correctly
2. **CORS Errors**: Verify the frontend URL is in the backend's CORS configuration
3. **Build Failures**: Check the build logs in Render dashboard
4. **Environment Variables**: Ensure all AWS credentials are set correctly

### Useful Commands:

- **View Logs**: Go to your service in Render dashboard → "Logs"
- **Redeploy**: Go to your service → "Manual Deploy"
- **Environment Variables**: Go to your service → "Environment"

## Cost

- **Free Tier**: 750 hours/month for each service
- **Database**: Free tier available
- **Custom Domains**: Available on paid plans

## Next Steps

1. Set up custom domains (optional)
2. Configure SSL certificates (automatic on Render)
3. Set up monitoring and alerts
4. Configure auto-scaling if needed

## Support

- [Render Documentation](https://render.com/docs)
- [Render Community](https://community.render.com)
- [Render Status](https://status.render.com) 