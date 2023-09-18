import KeyIcon from '@mui/icons-material/Key';
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import Typography from "@mui/material/Typography";
import { User } from "../hooks/Auth";

type CredentialsProps = {
  user: User;
};

export const Credentials = (props: CredentialsProps) => {
  const { user } = props;

  if (user.credentials.length === 0) {
    return (
      <Typography mb={4} variant="h4" component="h3" gutterBottom>
        passkey is not registered
      </Typography>
    );
  }
  return (
    <>
      <Typography mb={4} variant="h4" component="h3" gutterBottom>
          Registered passkeys
      </Typography>
      <List>
        {user.credentials.map((cred) => (
          <ListItem key={cred.name}>
            <ListItemIcon>
              <KeyIcon />
            </ListItemIcon>
            {cred.name}
          </ListItem>
        ))}
      </List>
    </>
  );
};
