import React from 'react'
import Paper from '@mui/material/Paper'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import Button from '@mui/material/Button'
import Stack from '@mui/material/Stack'

function Upload() {
  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '60vh' }}>
      <Paper elevation={2} sx={{ width: '100%', maxWidth: 400, p: 6, borderRadius: 4, boxShadow: 2 }}>
        <Typography variant="h4" fontWeight={700} gutterBottom align="left">
          Upload File
        </Typography>
        <Box component="form" sx={{ mt: 2 }}>
          <Stack spacing={3}>
            <Button variant="outlined" component="label" sx={{ borderRadius: 2, fontWeight: 600, py: 1.5 }}>
              Select File
              <input type="file" hidden />
            </Button>
            <Button variant="contained" color="primary" fullWidth sx={{ borderRadius: 2, fontWeight: 600, py: 1.5 }}>Upload</Button>
          </Stack>
        </Box>
      </Paper>
    </Box>
  )
}

export default Upload