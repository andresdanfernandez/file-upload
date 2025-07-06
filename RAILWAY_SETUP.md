# Railway Deployment Guide

## Overview
This is a file upload application with:
- **Backend**: Go API (Gin) with Supabase auth, AWS S3 storage, PostgreSQL
- **Frontend**: React (Vite) with Supabase auth, Material-UI

## Railway Setup

### 1. Create Two Services

#### Backend Service
1. **New Service** → **GitHub Repo**
2. **Repository**: `andresdanfernandez/file-upload`
3. **Root Directory**: `/` (leave empty for root)
4. **Name**: `backend` or `file-upload-api`

#### Frontend Service  
1. **New Service** → **GitHub Repo**
2. **Repository**: `andresdanfernandez/file-upload`
3. **Root Directory**: `frontend`
4. **Name**: `frontend` or `file-upload-frontend`

### 2. Backend Environment Variables

Add these to your backend service:

```
# Port
PORT=8080

# Database (use Railway PostgreSQL or external)
DB_HOST=your-db-host
DB_PORT=5432
DB_USER=your-db-user
DB_PASSWORD=your-db-password
DB_NAME=your-db-name

# AWS S3
AWS_REGION=us-east-1
AWS_ACCESS_KEY_ID=your-aws-access-key
AWS_SECRET_ACCESS_KEY=your-aws-secret-key
S3_BUCKET_NAME=your-s3-bucket-name

# Supabase
SUPABASE_URL=https://your-project.supabase.co
SUPABASE_ANON_KEY=your-supabase-anon-key
SUPABASE_SERVICE_ROLE_KEY=your-supabase-service-role-key
```

### 3. Frontend Environment Variables

Add these to your frontend service:

```
# API Base URL (your backend service URL)
VITE_API_BASE=https://your-backend-service.railway.app

# Supabase
VITE_SUPABASE_URL=https://your-project.supabase.co
VITE_SUPABASE_ANON_KEY=your-supabase-anon-key
```

### 4. Database Setup

#### Option A: Railway PostgreSQL
1. **New Service** → **Database** → **PostgreSQL**
2. Railway will auto-set these variables:
   ```
   PGHOST, PGPORT, PGUSER, PGPASSWORD, PGDATABASE, DATABASE_URL
   ```
3. Update your backend variables to use Railway's values:
   ```
   DB_HOST=${PGHOST}
   DB_PORT=${PGPORT}
   DB_USER=${PGUSER}
   DB_PASSWORD=${PGPASSWORD}
   DB_NAME=${PGDATABASE}
   ```

#### Option B: External Database
Use Supabase, AWS RDS, or any PostgreSQL provider.

### 5. AWS S3 Setup

1. **Create S3 Bucket** in AWS Console
2. **Create IAM User** with S3 access
3. **Get Access Keys** from IAM
4. **Configure CORS** on your S3 bucket:

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

1. **Create Project** at [supabase.com](https://supabase.com)
2. **Go to Settings** → **API**
3. **Copy Project URL and anon key**
4. **Create Service Role Key** for backend auth

### 7. Deployment Order

1. **Deploy Backend First**
   - Wait for it to be healthy
   - Test: `https://your-backend.railway.app/health`

2. **Set Frontend Variables**
   - Set `VITE_API_BASE` to your backend URL

3. **Deploy Frontend**
   - Wait for it to be healthy

### 8. Testing

1. **Visit your frontend URL**
2. **Sign up/sign in** with Supabase
3. **Test file upload/download**

## Troubleshooting

### Common Issues

**404 Errors:**
- Check both services are deployed and healthy
- Verify environment variables are set
- Check Railway logs

**CORS Errors:**
- Verify backend CORS includes frontend domain
- Check `VITE_API_BASE` is correct

**Auth Issues:**
- Verify Supabase environment variables
- Check database is initialized

**File Upload Issues:**
- Verify AWS S3 credentials
- Check S3 bucket permissions
- Verify bucket name is correct

### Manual Redeploy

If frontend doesn't auto-deploy:
1. Go to frontend service
2. **Deployments** tab
3. Click **"Deploy"**

## Environment Variables Reference

### Backend Required
```
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
```

### Frontend Required
```
VITE_API_BASE=https://your-backend.railway.app
VITE_SUPABASE_URL=https://your-project.supabase.co
VITE_SUPABASE_ANON_KEY=your-anon-key
``` 