import React from 'react'
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom'
import AppBar from '@mui/material/AppBar'
import Toolbar from '@mui/material/Toolbar'
import Button from '@mui/material/Button'
import CssBaseline from '@mui/material/CssBaseline'
import { ThemeProvider } from '@mui/material/styles'
import './App.css'

import FileList from './components/FileList'
import Upload from './components/Upload'
import Login from './components/Login'
import Home from './components/Home'
import Register from './components/Register'

import { theme } from './styles/Theme'

function App() {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Router>
        <AppBar position="static" color="primary" elevation={2} sx={{ borderRadius: 0, mb: 4 }}>
          <Toolbar sx={{ display: 'flex', gap: 2 }}>
            <Button color="inherit" component={Link} to="/" sx={{ borderRadius: 2, fontWeight: 600 }}>Home</Button>
            <Button color="inherit" component={Link} to="/login" sx={{ borderRadius: 2, fontWeight: 600 }}>Login</Button>
            <Button color="inherit" component={Link} to="/upload" sx={{ borderRadius: 2, fontWeight: 600 }}>Upload</Button>
            <Button color="inherit" component={Link} to="/files" sx={{ borderRadius: 2, fontWeight: 600 }}>File List</Button>
            <Button color="inherit" component={Link} to="/register" sx={{ borderRadius: 2, fontWeight: 600 }}>Register</Button>
          </Toolbar>
        </AppBar>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/upload" element={<Upload />} />
          <Route path="/files" element={<FileList />} />
          <Route path="/register" element={<Register />} />
        </Routes>
      </Router>
    </ThemeProvider>
  )
}

export default App
