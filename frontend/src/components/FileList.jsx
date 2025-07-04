import React from 'react'
import Paper from '@mui/material/Paper'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import Table from '@mui/material/Table'
import TableBody from '@mui/material/TableBody'
import TableCell from '@mui/material/TableCell'
import TableContainer from '@mui/material/TableContainer'
import TableHead from '@mui/material/TableHead'
import TableRow from '@mui/material/TableRow'
import Button from '@mui/material/Button'

function FileList() {
  const files = [
    { name: 'file1.txt', size: '2 KB', uploaded: '2024-06-01' },
    { name: 'photo.jpg', size: '1.5 MB', uploaded: '2024-06-02' },
  ]
  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '60vh' }}>
      <Paper elevation={2} sx={{ width: '100%', maxWidth: 700, p: 6, borderRadius: 4, boxShadow: 2 }}>
        <Typography variant="h4" fontWeight={700} gutterBottom align="left">
          Your Files
        </Typography>
        <TableContainer component={Paper} sx={{ borderRadius: 2, boxShadow: 0, mt: 2 }}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell sx={{ fontWeight: 700 }}>Name</TableCell>
                <TableCell sx={{ fontWeight: 700 }}>Size</TableCell>
                <TableCell sx={{ fontWeight: 700 }}>Uploaded</TableCell>
                <TableCell sx={{ fontWeight: 700 }}>Action</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {files.map((file, idx) => (
                <TableRow key={idx} hover>
                  <TableCell>{file.name}</TableCell>
                  <TableCell>{file.size}</TableCell>
                  <TableCell>{file.uploaded}</TableCell>
                  <TableCell>
                    <Button size="small" variant="outlined" sx={{ borderRadius: 2, fontWeight: 600 }}>Download</Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Paper>
    </Box>
  )
}

export default FileList