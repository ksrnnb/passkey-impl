import * as React from 'react';
import CssBaseline from '@mui/material/CssBaseline';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { Button } from '@mui/material';
import * as client from "../httpClient/client";
import { useAuth } from '../hooks/Auth';
import { useNavigate } from 'react-router-dom';

const defaultTheme = createTheme();

export default function StickyFooter() {
  const { unsetToken } = useAuth();
  const navigate = useNavigate();

  const  handleSignOut = async () => {
    await client.post("/signout");
    unsetToken();

    // TODO: investigate why navigate cannot work well
    //       maybe state is not updated here...
    // navigate("/signin");
    window.location.href = "/signin";
  };

  return (
    <ThemeProvider theme={defaultTheme}>
      <Box
        sx={{
          display: 'flex',
          flexDirection: 'column',
          minHeight: '100vh',
        }}
      >
        <CssBaseline />
        <Container component="main" sx={{ mt: 8, mb: 2 }} maxWidth="sm">
          <Typography variant="h2" component="h1" gutterBottom>
            Passkey sample
          </Typography>
          <Typography mb={4} variant="h5" component="h2" gutterBottom>
            {'Pin a footer to the bottom of the viewport.'}
            {'The footer will move as the main element of the page grows.'}
          </Typography>
          <Button variant="contained" onClick={handleSignOut}>
            Sign Out
          </Button>
        </Container>
      </Box>
    </ThemeProvider>
  );
}
