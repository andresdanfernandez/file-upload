import { Container, Paper } from '@mui/material'

export default function PageWrapper({ children }) {
    return (
      <Container maxWidth="sm" sx={{ mt: 6, mb: 6 }}>
        <Paper elevation={3} sx={{ p: 4, borderRadius: 4, minHeight: 300 }}>
          {children}
        </Paper>
      </Container>
    )
  }