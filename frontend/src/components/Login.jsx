import React from 'react'
import Paper from '@mui/material/Paper'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import TextField from '@mui/material/TextField'
import Button from '@mui/material/Button'
import Stack from '@mui/material/Stack'

function Login() {
  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '60vh' }}>
      <Paper elevation={2} sx={{ width: '100%', maxWidth: 400, p: 6, borderRadius: 4, boxShadow: 2 }}>
        <Typography variant="h4" fontWeight={700} gutterBottom align="left">
          Login
        </Typography>
        <Box component="form" sx={{ mt: 2 }}>
          <Stack spacing={3}>
            <TextField label="Username" fullWidth variant="outlined" InputProps={{ sx: { borderRadius: 2 } }} />
            <TextField label="Password" type="password" fullWidth variant="outlined" InputProps={{ sx: { borderRadius: 2 } }} />
            <Button variant="contained" color="primary" fullWidth sx={{ borderRadius: 2, fontWeight: 600, py: 1.5 }}>Login</Button>
          </Stack>
        </Box>
      </Paper>
    </Box>
  )
}

export default Login