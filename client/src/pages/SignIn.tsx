import * as React from 'react';
import { useNavigate } from 'react-router-dom';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';

import authContext from '../context/AuthContext';
import * as client from "../httpClient/client";
import { toArrayBuffer, toBase64Url } from '../utils/array_buffer';

const defaultTheme = createTheme();

type SignInResponse = {
  token: string;
};

type StartSingInWithPasskeyResponse = {
  publicKey: {
    challenge: string;
    rpId: string;
    timeout: number;
    userVerification: "discouraged" | "preferred" | "required";
  }
};

type SignInWithPasskeyRequest = {
  id: string;
  rawId: string;
  type: string;
  authenticatorAttachment: string;
  response: {
    clientDataJSON: string;
    authenticatorData: string;
    signature: string;
    userHandle: string;
  },
};

type SignInWithPasskeyResponse = {
  token?: string;
}

const credentialToSignInWithPasskeyRequest = (cred: Credential): SignInWithPasskeyRequest => {
  const pubKeyCred = cred as PublicKeyCredential;
  const aar = pubKeyCred.response as AuthenticatorAssertionResponse;
  const req: SignInWithPasskeyRequest = {
    id: pubKeyCred.id,
    rawId: pubKeyCred.id,
    authenticatorAttachment: pubKeyCred.authenticatorAttachment ?? "",
    response: {
      authenticatorData: toBase64Url(aar.authenticatorData),
      clientDataJSON: toBase64Url(aar.clientDataJSON),
      signature: toBase64Url(aar.signature),
      userHandle: aar.userHandle ? toBase64Url(aar.userHandle) : "",
    },
    type: pubKeyCred.type,
  };

  return req;
};

export default function SignIn() {
  const { setToken } = React.useContext(authContext);
  const navigate = useNavigate();

  const handleSignIn = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    const userId = data.get('userId');
    const password = data.get('password');

    const res: SignInResponse = await client.Post("/signin", {userId, password})
                                              .then(res => res.json());

    setToken(res.token);

    // TODO: use navigate
    //       if navigate is used, request of navigator.credentials.get will be pending.
    //       so, passkey registration cannot start in next page...
    // navigate("/");
    window.location.href = "/";
  };

  const signInWithPasskey = React.useCallback(async () => {
    const res: StartSingInWithPasskeyResponse = await client.Post("/signin/passkey/start").then(res => res.json());
  
    const options: CredentialRequestOptions = {
      publicKey: {
        challenge: toArrayBuffer(res.publicKey.challenge),
        rpId: res.publicKey.rpId,
        timeout: res.publicKey.timeout,
        userVerification: res.publicKey.userVerification,
        // not specify allowCredential to use auto fill
      },
      mediation: 'conditional',
    }

    const cred = await navigator.credentials.get(options);
    if (!cred) {
      return;
    }
  
    const req = credentialToSignInWithPasskeyRequest(cred);
  
    const signInRes: SignInWithPasskeyResponse = await client.Post("/signin/passkey", req).then(res => res.json());
    
    if (signInRes.token) {
      setToken(signInRes.token);
    }
  }, [setToken]);
  
  const executeSignInWithPasskey = React.useCallback(async () => {
    if (PublicKeyCredential && PublicKeyCredential.isConditionalMediationAvailable) {
      const isCMA = await PublicKeyCredential.isConditionalMediationAvailable();
      if (isCMA) {
        // if conditional UI is available, execute sign in with passkey
        signInWithPasskey();
      }
    }
  }, [signInWithPasskey]);
  
  React.useEffect(() => {
    // TODO: don't use useEffect
    //       because in StrictMode, useEffect is executed twice but this executeSignInWithPasskey doesn't expect to call twice.
    executeSignInWithPasskey();
  }, [executeSignInWithPasskey]);

  return (
    <ThemeProvider theme={defaultTheme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <Typography component="h1" variant="h5">
            Sign in
          </Typography>
          <Box component="form" onSubmit={handleSignIn} noValidate sx={{ mt: 1 }}>
            <TextField
              margin="normal"
              required
              fullWidth
              id="userId"
              label="User Id"
              name="userId"
              autoComplete="userId webauthn"
              defaultValue="sample"
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="current-password"
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Sign In
            </Button>
            {/* <Grid container>
              <Grid item>
                <Link href="/signup" variant="body2">
                  {"Don't have an account? Sign Up"}
                </Link>
              </Grid>
            </Grid> */}
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
}
