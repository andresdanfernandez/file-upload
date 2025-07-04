import { Container, Paper, Typography, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Button } from '@mui/material'

export default function FileList() {
  const files = [
    { name: 'file1.txt', size: '2 KB', uploaded: '2024-06-01' },
    { name: 'photo.jpg', size: '1.5 MB', uploaded: '2024-06-02' },
  ]
  return (
    <Container maxWidth="md" sx={{ mt: 6, mb: 6 }}>
      <Paper elevation={3} sx={{ p: 4, borderRadius: 4 }}>
        <Typography variant="h4" gutterBottom fontWeight={700}>Your Files</Typography>
        <TableContainer component={Paper} sx={{ borderRadius: 3, boxShadow: 0, mt: 2 }}>
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
                <TableRow key={idx} hover sx={{ borderRadius: 2 }}>
                  <TableCell>{file.name}</TableCell>
                  <TableCell>{file.size}</TableCell>
                  <TableCell>{file.uploaded}</TableCell>
                  <TableCell>
                    <Button size="small" variant="contained" sx={{ borderRadius: 2, fontWeight: 600 }}>Download</Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Paper>
    </Container>
  )
}