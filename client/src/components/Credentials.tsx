import KeyIcon from '@mui/icons-material/Key';
import Box from "@mui/material/Box";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import Typography from "@mui/material/Typography";
import { User } from "../hooks/Auth";
import { CredentialDeleteIcon, CredentialDeleteIconProps } from './CredentialDeleteIcon';

type CredentialsProps = {
  user: User;
  deleteCredential: CredentialDeleteIconProps["deleteCredential"];
};

export const Credentials = (props: CredentialsProps) => {
  const { user, deleteCredential } = props;

  if (user.credentials.length === 0) {
    return (
      <Typography mb={4} variant="h4" component="h3" gutterBottom>
        passkey is not registered
      </Typography>
    );
  }
  return (
    <Box mb={5}>
      <Typography mb={4} variant="h4" component="h3" gutterBottom>
          Registered passkeys
      </Typography>
      <List>
        {user.credentials.map((cred) => (
          <ListItem key={cred.id} secondaryAction={<CredentialDeleteIcon credentialId={cred.id} deleteCredential={deleteCredential}/>}>
            <ListItemIcon>
              <KeyIcon />
            </ListItemIcon>
            {cred.name}
          </ListItem>
        ))}
      </List>
    </Box>
  );
};
