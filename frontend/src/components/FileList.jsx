import React, { useEffect, useState } from 'react'
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
import Grow from '@mui/material/Grow'
import Fade from '@mui/material/Fade'
import { listFiles, getToken, getDownloadUrl } from '../api'

function FileList() {
  const [files, setFiles] = useState([]);
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchFiles = async () => {
      setLoading(true);
      setError('');
      try {
        const token = getToken();
        if (!token) {
          setError('Not authenticated');
          setLoading(false);
          return;
        }
        const res = await listFiles(token);
        if (Array.isArray(res)) {
          setFiles(res);
        } else {
          setError(res.error || 'Failed to fetch files');
        }
      } catch (err) {
        setError('Failed to fetch files');
      } finally {
        setLoading(false);
      }
    };
    fetchFiles();
  }, []);

  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '60vh' }}>
      <Grow in={true} timeout={500}>
        <Paper elevation={2} sx={{ width: '100%', maxWidth: 700, p: 6, borderRadius: 4, boxShadow: 2 }}>
          <Fade in={true} timeout={900}>
            <Box>
              <Typography variant="h4" fontWeight={700} gutterBottom align="left">
                Your Files
              </Typography>
              {loading ? <div>Loading...</div> : error ? <div style={{color:'red'}}>{error}</div> : (
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
              )}
            </Box>
          </Fade>
        </Paper>
      </Grow>
    </Box>
  )
}

export default FileList