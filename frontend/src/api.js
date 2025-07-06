// API utility for backend interaction
// Supports: upload, list files, download
// Handles Supabase session token and sending

import { supabase } from './supabase';

const API_BASE = import.meta.env.VITE_API_BASE || 'http://localhost:8080';

// --- Supabase Token Helpers ---
export async function getSupabaseToken() {
  const { data: { session } } = await supabase.auth.getSession();
  console.log('Supabase session:', session ? 'exists' : 'null');
  console.log('Access token:', session?.access_token ? 'exists' : 'null');
  
  if (session?.access_token) {
    console.log('Token preview:', session.access_token.substring(0, 20) + '...');
    console.log('Token length:', session.access_token.length);
  }
  
  return session?.access_token;
}

// file upload
export async function uploadFile(file) {
  const token = await getSupabaseToken();
  if (!token) {
    throw new Error('Not authenticated');
  }
  
  console.log('Upload: Sending token length:', token.length);
  
  const formData = new FormData();
  formData.append('file', file);
  const res = await fetch(`${API_BASE}/upload`, {
    method: 'POST',
    headers: { 'Authorization': `Bearer ${token}` },
    body: formData
  });
  
  if (!res.ok) {
    const errorText = await res.text();
    console.error('Upload failed:', res.status, errorText);
    throw new Error(`Upload failed: ${res.status} - ${errorText}`);
  }
  
  return res.json();
}

// list files
export async function listFiles() {
  const token = await getSupabaseToken();
  if (!token) {
    throw new Error('Not authenticated');
  }
  
  console.log('List files: Sending token length:', token.length);
  
  const res = await fetch(`${API_BASE}/files`, {
    headers: { 'Authorization': `Bearer ${token}` }
  });
  
  if (!res.ok) {
    const errorText = await res.text();
    console.error('List files failed:', res.status, errorText);
    throw new Error(`List files failed: ${res.status} - ${errorText}`);
  }
  
  return res.json();
}

// download file 
export async function downloadFile(key) {
  const token = await getSupabaseToken();
  if (!token) {
    throw new Error('Not authenticated');
  }
  
  console.log('Download: Sending token length:', token.length);
  
  const res = await fetch(`${API_BASE}/download/${encodeURIComponent(key)}`, {
    headers: { 'Authorization': `Bearer ${token}` }
  });
  
  if (!res.ok) {
    const error = await res.json();
    throw new Error(error.error || 'Download failed');
  }
  
  // Get the presigned URL from the response
  const data = await res.json();
  return data.url;
}