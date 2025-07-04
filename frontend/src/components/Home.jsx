import React from 'react'
import Paper from '@mui/material/Paper'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'

function Home() {
  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '60vh' }}>
      <Paper elevation={2} sx={{ width: '100%', maxWidth: 600, p: 6, borderRadius: 4, boxShadow: 2 }}>
        <Typography variant="h3" fontWeight={700} gutterBottom align="left">
          File Upload App
        </Typography>
        <Typography variant="h6" color="text.secondary" align="left" sx={{ mt: 2 }}>
          Use the navigation above to login, upload, or view your files.
        </Typography>
      </Paper>
    </Box>
  )
}

export default Home