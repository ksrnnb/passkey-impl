import * as React from 'react';
import CssBaseline from '@mui/material/CssBaseline';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import Button from '@mui/material/Button';
import * as client from "../httpClient/client";
import { useAuth } from '../hooks/Auth';
import { Credentials } from '../components/Credentials';

const defaultTheme = createTheme();

function toBase64Url(buffer: ArrayBuffer) {
  const base64 = window.btoa(String.fromCharCode(...new Uint8Array(buffer)));
  return base64.replace(/=/g, "").replace(/\+/g, "-").replace(/\//g, "_");
}

function toArrayBuffer(str: string) {
  const base64 = str.replace(/-/g, "+").replace(/_/g, "/");
  const binStr = window.atob(base64);
  const bin = new Uint8Array(binStr.length);
  for (let i = 0; i < binStr.length; i++) {
    bin[i] = binStr.charCodeAt(i);
  }
  return bin.buffer;
}


type PubKeyCredParam = {
  type: "public-key",
  alg: number,
};

type StartRegistrationResponse = {
  publicKey: {
    rp: {
      id: string,
      name: string,
    },
    user: {
      id: string,
      name: string,
      displayName: string,
    },
    pubKeyCredParams: PubKeyCredParam[],
    challenge: string,
    authenticatorSelection: {
      requireResidentKey: boolean,
      userVerification: "required" | "preferrer" | "discouraged";
    },
    timeout: number,
  },
}

type RegisterPasskeyRequest = {
  id: string,
  type: string,
  authenticatorAttachment: string,
  response: {
    clientDataJSON: string,
    attestationObject: string,
    transports: string[],
  },
}

export default function StickyFooter() {
  const { user, updateUser, unsetToken } = useAuth();

  const  handleSignOut = async () => {
    await client.post("/signout");
    unsetToken();

    // TODO: investigate why navigate cannot work well
    //       maybe state is not updated here...
    // navigate("/signin");
    window.location.href = "/signin";
  };

  const handleRegisterPasskey = async () => {
    const res: StartRegistrationResponse = await client.post("/passkey/register/start").then(res => res.json());
    console.log("challenge:", res.publicKey.challenge)

    const options: CredentialCreationOptions = {
      publicKey: {
        challenge: toArrayBuffer(res.publicKey.challenge),
        rp: {
          // An RP ID is based on a host's domain name.
          // It does not itself include a scheme or port, as an origin does.
          // ref: https://www.w3.org/TR/webauthn-2/#rp-id
          name: res.publicKey.rp.name,
        },
        user: {
          id: toArrayBuffer(res.publicKey.user.id),
          name: res.publicKey.user.name,
          displayName: res.publicKey.user.displayName ?? res.publicKey.user.name,
        },
        pubKeyCredParams: res.publicKey.pubKeyCredParams,
        timeout: res.publicKey.timeout,
      },
    };
  
    const cred = await navigator.credentials.create(options)
    .catch(err => {
      // do nothing
      // when user cancel registration, error will be thrown
    });
  
    if (!cred) {
      return;
    }
    const pubKeyCred = cred as PublicKeyCredential;
    const response = pubKeyCred.response as AuthenticatorAttestationResponse;
    const req: RegisterPasskeyRequest = {
      id: pubKeyCred.id,
      type: pubKeyCred.type,
      authenticatorAttachment: pubKeyCred.authenticatorAttachment ?? "",
      response: {
        clientDataJSON: toBase64Url(response.clientDataJSON),
        attestationObject: toBase64Url(response.attestationObject),
        transports: response.getTransports(),
      },
    }
    
    const registerRes = await client.post("/passkey/register", req).then(res => res.json());

    console.log(registerRes);
    updateUser();
  }

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
          <Typography mb={4} variant="h4" component="h2" gutterBottom>
            Passkey setting
          </Typography>
          <Box mb={5} pb={5}>
            {user ? <Credentials user={user}/> : <></>}
            {user?.credentials.length === 0 && (
              <Button variant="contained" onClick={handleRegisterPasskey}>
                Register passkey
              </Button>
            )}
          </Box>
          <Box>
            <Button variant="outlined" color="secondary" onClick={handleSignOut}>
              Sign Out
            </Button>
          </Box>
        </Container>
      </Box>
    </ThemeProvider>
  );
}
