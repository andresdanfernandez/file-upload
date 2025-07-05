import React, { useState } from 'react'
import Paper from '@mui/material/Paper'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import Button from '@mui/material/Button'
import Stack from '@mui/material/Stack'
import Grow from '@mui/material/Grow'
import Fade from '@mui/material/Fade'
import { uploadFile } from '../api'

function Upload() {
  const [file, setFile] = useState(null)
  const [result, setResult] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const handleUpload = async (e) => {
    e.preventDefault()
    setLoading(true)
    setResult('')
    setError('')
    try {
      const res = await uploadFile(file)
      if (res.url) {
        setResult(`File uploaded! URL: ${res.url}`)
      } else {
        setError(res.error || 'Upload failed')
      }
    } catch (err) {
      setError(err.message || 'Upload failed')
    } finally {
      setLoading(false)
    }
  }

  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '60vh' }}>
      <Grow in={true} timeout={500}>
        <Paper elevation={2} sx={{ width: '100%', maxWidth: 400, p: 6, borderRadius: 4, boxShadow: 2 }}>
          <Fade in={true} timeout={900}>
            <Box>
              <Typography variant="h4" fontWeight={700} gutterBottom align="left">
                Upload File
              </Typography>
              <Box component="form" sx={{ mt: 2 }} onSubmit={handleUpload}>
                <Stack spacing={3}>
                  <Button variant="outlined" component="label" sx={{ borderRadius: 2, fontWeight: 600, py: 1.5 }}>
                    Select File
                    <input type="file" hidden onChange={e => setFile(e.target.files[0])} required />
                  </Button>
                  <Button variant="contained" color="primary" fullWidth sx={{ borderRadius: 2, fontWeight: 600, py: 1.5 }} type="submit" disabled={loading || !file}>
                    {loading ? 'Uploading...' : 'Upload'}
                  </Button>
                </Stack>
                {result && <div style={{ color: 'green' }}>{result}</div>}
                {error && <div style={{ color: 'red' }}>{error}</div>}
              </Box>
            </Box>
          </Fade>
        </Paper>
      </Grow>
    </Box>
  )
}

export default Upload