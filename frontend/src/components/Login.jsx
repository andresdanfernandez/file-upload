import React, { useState } from 'react'
import Paper from '@mui/material/Paper'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import TextField from '@mui/material/TextField'
import Button from '@mui/material/Button'
import Stack from '@mui/material/Stack'
import Grow from '@mui/material/Grow'
import Fade from '@mui/material/Fade'
import { login, saveToken } from '../api'

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');
    try {
      const res = await login(email, password);
      if (res.token) {
        saveToken(res.token);
        // Optionally redirect or update parent state
        window.location.reload();
      } else {
        setError(res.error || 'Login failed');
      }
    } catch (err) {
      setError('Login failed');
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
                Login
              </Typography>
              <Box component="form" sx={{ mt: 2 }} onSubmit={handleSubmit}>
                <Stack spacing={3}>
                  <TextField label="Email" fullWidth variant="outlined" InputProps={{ sx: { borderRadius: 2 } }} value={email} onChange={(e) => setEmail(e.target.value)} />
                  <TextField label="Password" type="password" fullWidth variant="outlined" InputProps={{ sx: { borderRadius: 2 } }} value={password} onChange={(e) => setPassword(e.target.value)} />
                  <Button variant="contained" color="primary" fullWidth sx={{ borderRadius: 2, fontWeight: 600, py: 1.5 }} type="submit" disabled={loading}>{loading ? 'Logging in...' : 'Login'}</Button>
                </Stack>
                {error && <div style={{ color: 'red' }}>{error}</div>}
              </Box>
            </Box>
          </Fade>
        </Paper>
      </Grow>
    </Box>
  )
}

export default Login