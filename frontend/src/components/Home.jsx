import React from 'react'
import Paper from '@mui/material/Paper'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import Grow from '@mui/material/Grow'
import Fade from '@mui/material/Fade'

function Home() {
  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '60vh' }}>
      <Grow in={true} timeout={500}>
        <Paper elevation={2} sx={{ width: '100%', maxWidth: 600, p: 6, borderRadius: 4, boxShadow: 2 }}>
          <Fade in={true} timeout={900}>
            <Box>
              <Typography variant="h3" fontWeight={700} gutterBottom align="left">
                File Upload App
              </Typography>
              <Typography variant="h6" color="text.secondary" align="left" sx={{ mt: 2 }}>
                Use the navigation above to login, upload, or view your files.
              </Typography>
            </Box>
          </Fade>
        </Paper>
      </Grow>
    </Box>
  )
}

export default Home