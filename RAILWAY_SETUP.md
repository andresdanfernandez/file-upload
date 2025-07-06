# Railway Deployment Setup Guide

## Overview
This project consists of two services:
1. **Backend** (Go API) - Handles file uploads, authentication, and file management
2. **Frontend** (React) - User interface for file management

## Railway Project Setup

### 1. Create Two Services in Railway

Create two separate services in your Railway project:
- **Backend Service** (for the Go API)
- **Frontend Service** (for the React app)

### 2. Backend Service Configuration

**Repository**: Point to your main repository
**Root Directory**: `/` (root of the repository)
**Build Command**: Uses the Dockerfile in the root directory

**Required Environment Variables**:
```
PORT=8080
DB_HOST=your-database-host
DB_PORT=5432
DB_USER=your-database-user
DB_PASSWORD=your-database-password
DB_NAME=your-database-name
AWS_REGION=your-aws-region
AWS_ACCESS_KEY_ID=your-aws-access-key
AWS_SECRET_ACCESS_KEY=your-aws-secret-key
S3_BUCKET_NAME=your-s3-bucket-name
SUPABASE_URL=your-supabase-project-url
SUPABASE_ANON_KEY=your-supabase-anon-key
SUPABASE_SERVICE_ROLE_KEY=your-supabase-service-role-key
JWT_SECRET=your-jwt-secret
```

### 3. Frontend Service Configuration

**Repository**: Point to your main repository
**Root Directory**: `/frontend`
**Build Command**: Uses the Dockerfile in the frontend directory

**Required Environment Variables**:
```
PORT=4173
VITE_API_BASE=https://your-backend-service-url.railway.app
VITE_SUPABASE_URL=your-supabase-project-url
VITE_SUPABASE_ANON_KEY=your-supabase-anon-key
```

### 4. Database Setup

You'll need a PostgreSQL database. You can:
- Use Railway's PostgreSQL plugin
- Use an external database (Supabase, AWS RDS, etc.)

If using Railway's PostgreSQL:
1. Add PostgreSQL plugin to your project
2. Use the connection details provided by Railway

### 5. S3 Setup

You'll need an AWS S3 bucket for file storage:
1. Create an S3 bucket in AWS
2. Create an IAM user with S3 access
3. Get the access key and secret key
4. Configure CORS on your S3 bucket:

```json
[
    {
        "AllowedHeaders": ["*"],
        "AllowedMethods": ["GET", "PUT", "POST", "DELETE"],
        "AllowedOrigins": ["*"],
        "ExposeHeaders": []
    }
]
```

### 6. Supabase Setup

1. Create a Supabase project
2. Get your project URL and anon key
3. Create a service role key for backend authentication
4. Set up your database schema (see `scripts/init_db.sql`)

### 7. Deployment Steps

1. **Deploy Backend First**:
   - Push your code to GitHub
   - Deploy the backend service
   - Wait for it to be healthy

2. **Deploy Frontend**:
   - Set the `VITE_API_BASE` to your backend service URL
   - Deploy the frontend service
   - Wait for it to be healthy

3. **Test the Deployment**:
   - Visit your frontend service URL
   - Test authentication
   - Test file upload/download

### 8. Troubleshooting

**404 Errors**:
- Check that both services are deployed and healthy
- Verify environment variables are set correctly
- Check Railway logs for any build or runtime errors

**CORS Errors**:
- Verify the backend CORS configuration includes your frontend domain
- Check that the frontend is using the correct backend URL

**Authentication Issues**:
- Verify Supabase environment variables are correct
- Check that the database is properly initialized
- Verify JWT secret is set

**File Upload Issues**:
- Verify AWS S3 credentials are correct
- Check S3 bucket permissions
- Verify the bucket name is correct

### 9. Monitoring

- Use Railway's built-in logging to monitor both services
- Set up alerts for service health
- Monitor database and S3 usage

## Environment Variables Reference

### Backend (.env)
```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=fileupload
AWS_REGION=us-east-1
AWS_ACCESS_KEY_ID=your-key
AWS_SECRET_ACCESS_KEY=your-secret
S3_BUCKET_NAME=your-bucket
SUPABASE_URL=https://your-project.supabase.co
SUPABASE_ANON_KEY=your-anon-key
SUPABASE_SERVICE_ROLE_KEY=your-service-role-key
JWT_SECRET=your-jwt-secret
```

### Frontend (.env)
```env
VITE_API_BASE=http://localhost:8080
VITE_SUPABASE_URL=https://your-project.supabase.co
VITE_SUPABASE_ANON_KEY=your-anon-key
``` 