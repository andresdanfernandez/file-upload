import { Typography, Button, Box, Stack } from '@mui/material'
import PageWrapper from './PageWrapper'

export default function Upload() {
    return (
      <PageWrapper>
        <Typography variant="h4" gutterBottom fontWeight={700}>Upload File</Typography>
        <Box component="form" sx={{ mt: 2 }}>
          <Stack spacing={3}>
            <Button variant="outlined" component="label" sx={{ borderRadius: 2, py: 1.5, fontWeight: 600 }}>
              Select File
              <input type="file" hidden />
            </Button>
            <Button variant="contained" color="primary" fullWidth sx={{ borderRadius: 2, py: 1.5, fontWeight: 600 }}>Upload</Button>
          </Stack>
        </Box>
      </PageWrapper>
    )
  }