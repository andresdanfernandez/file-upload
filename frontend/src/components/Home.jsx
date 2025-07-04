import { Typography } from '@mui/material'
import PageWrapper from './PageWrapper'

export default function Home() {
    return (
      <PageWrapper>
        <Typography variant="h4" gutterBottom fontWeight={700}>Welcome to File Upload App</Typography>
        <Typography color="text.secondary">Use the navigation above to login, upload, or view your files.</Typography>
      </PageWrapper>
    ) 
}