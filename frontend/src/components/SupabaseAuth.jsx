import React, { useState, useEffect } from 'react'
import { supabase } from '../supabase'
import { Auth } from '@supabase/auth-ui-react'
import { ThemeSupa } from '@supabase/auth-ui-shared'
import Box from '@mui/material/Box'
import Paper from '@mui/material/Paper'
import Typography from '@mui/material/Typography'

function SupabaseAuth() {
  const [session, setSession] = useState(null)

  useEffect(() => {
    supabase.auth.getSession().then(({ data: { session } }) => {
      setSession(session)
    })

    const {
      data: { subscription },
    } = supabase.auth.onAuthStateChange((_event, session) => {
      setSession(session)
    })

    return () => subscription.unsubscribe()
  }, [])

  if (!session) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '60vh' }}>
        <Paper elevation={2} sx={{ width: '100%', maxWidth: 400, p: 4, borderRadius: 4, boxShadow: 2 }}>
          <Typography variant="h4" fontWeight={700} gutterBottom align="center" sx={{ mb: 3 }}>
            Welcome to File Upload
          </Typography>
          <Auth
            supabaseClient={supabase}
            appearance={{ theme: ThemeSupa }}
            providers={['google']}
            theme="dark"
          />
        </Paper>
      </Box>
    )
  }

  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '60vh' }}>
      <Paper elevation={2} sx={{ width: '100%', maxWidth: 400, p: 4, borderRadius: 4, boxShadow: 2 }}>
        <Typography variant="h4" fontWeight={700} gutterBottom align="center">
          Welcome, {session.user.email}!
        </Typography>
        <Typography variant="body1" align="center" sx={{ mb: 3 }}>
          You are now authenticated with Supabase.
        </Typography>
        <button 
          onClick={() => supabase.auth.signOut()}
          style={{
            width: '100%',
            padding: '12px',
            backgroundColor: '#f44336',
            color: 'white',
            border: 'none',
            borderRadius: '8px',
            cursor: 'pointer',
            fontWeight: 600
          }}
        >
          Sign Out
        </button>
      </Paper>
    </Box>
  )
}

export default SupabaseAuth 