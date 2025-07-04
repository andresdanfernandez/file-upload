import { Typography, Button, Box, Stack, TextField } from '@mui/material'
import PageWrapper from './PageWrapper'

export default function Login() {
    return (
      <PageWrapper>
        <Typography variant="h4" gutterBottom fontWeight={700}>Login</Typography>
        <Box component="form" sx={{ mt: 2 }}>
          <Stack spacing={3}>
            <TextField label="Username" fullWidth variant="outlined" InputProps={{ sx: { borderRadius: 2 } }} />
            <TextField label="Password" type="password" fullWidth variant="outlined" InputProps={{ sx: { borderRadius: 2 } }} />
            <Button variant="contained" color="primary" fullWidth sx={{ borderRadius: 2, py: 1.5, fontWeight: 600 }}>Login</Button>
          </Stack>
        </Box>
      </PageWrapper>
    )
  }