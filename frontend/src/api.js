// API utility for backend interaction
// Supports: register, login, upload, list files, download
// Handles JWT storage and sending
// Extensible for third-party auth (see comments below)

const API_BASE = import.meta.env.VITE_API_BASE || 'http://localhost:8080';

// --- JWT Helpers ---
export function saveToken(token) {
  localStorage.setItem('jwt_token', token);
}

export function getToken() {
  return localStorage.getItem('jwt_token');
}

export function clearToken() {
  localStorage.removeItem('jwt_token');
}

// auth
export async function register(email, password) {
  const res = await fetch(`${API_BASE}/register`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  });
  return res.json();
}

export async function login(email, password) {
  const res = await fetch(`${API_BASE}/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  });
  return res.json();
}

// file upload
export async function uploadFile(file, token) {
  const formData = new FormData();
  formData.append('file', file);
  const res = await fetch(`${API_BASE}/upload`, {
    method: 'POST',
    headers: { 'Authorization': `Bearer ${token}` },
    body: formData
  });
  return res.json();
}

// list files
export async function listFiles(token) {
  const res = await fetch(`${API_BASE}/files`, {
    headers: { 'Authorization': `Bearer ${token}` }
  });
  return res.json();
}

// download file 
export function getDownloadUrl(key) {
  // This returns the backend download endpoint, which will redirect to the file
  return `${API_BASE}/download/${encodeURIComponent(key)}`;
}