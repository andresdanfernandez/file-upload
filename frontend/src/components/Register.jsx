import React, { useState } from 'react';
import Paper from '@mui/material/Paper';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import Grow from '@mui/material/Grow';
import Fade from '@mui/material/Fade';
import { useNavigate } from 'react-router-dom';

const Register = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');
    setLoading(true);
    try {
      const res = await fetch('/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      });
      const data = await res.json();
      if (res.ok) {
        setSuccess('Registration successful! Redirecting to login...');
        setTimeout(() => navigate('/login'), 1500);
      } else {
        setError(data.error || 'Registration failed');
      }
    } catch (err) {
      setError('Network error');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '60vh' }}>
      <Grow in={true} timeout={500}>
        <Paper elevation={2} sx={{ width: '100%', maxWidth: 400, p: 6, borderRadius: 4, boxShadow: 2 }}>
          <Fade in={true} timeout={900}>
            <Box>
              <Typography variant="h4" fontWeight={700} gutterBottom align="left">
                Register
              </Typography>
              <Box component="form" sx={{ mt: 2 }} onSubmit={handleSubmit}>
                <Stack spacing={3}>
                  <TextField label="Email" fullWidth variant="outlined" InputProps={{ sx: { borderRadius: 2 } }} value={email} onChange={e => setEmail(e.target.value)} required />
                  <TextField label="Password" type="password" fullWidth variant="outlined" InputProps={{ sx: { borderRadius: 2 } }} value={password} onChange={e => setPassword(e.target.value)} required minLength={6} />
                  <Button variant="contained" color="primary" fullWidth sx={{ borderRadius: 2, fontWeight: 600, py: 1.5 }} type="submit" disabled={loading}>{loading ? 'Registering...' : 'Register'}</Button>
                </Stack>
                {error && <div style={{ color: 'red', marginTop: 12 }}>{error}</div>}
                {success && <div style={{ color: 'green', marginTop: 12 }}>{success}</div>}
              </Box>
            </Box>
          </Fade>
        </Paper>
      </Grow>
    </Box>
  );
};

export default Register; 