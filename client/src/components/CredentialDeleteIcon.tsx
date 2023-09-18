import DeleteIcon from '@mui/icons-material/Delete';
import  IconButton  from '@mui/material/IconButton';

export type CredentialDeleteIconProps = {
  credentialId: string;
  deleteCredential: (credentialId: string) => void;
};

export const CredentialDeleteIcon = (props: CredentialDeleteIconProps) => {
  const { credentialId, deleteCredential } = props;

  return (
    <IconButton area-label="delete" onClick={() => deleteCredential(credentialId)}>
      <DeleteIcon />
    </IconButton>
  );
};
