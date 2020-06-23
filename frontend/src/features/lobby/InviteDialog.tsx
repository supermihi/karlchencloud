import * as React from 'react';
import Dialog from '@material-ui/core/Dialog';
import DialogTitle from '@material-ui/core/DialogTitle';
import Input from '@material-ui/core/Input';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import IconButton from '@material-ui/core/IconButton';
import AssignmentOutlinedIcon from '@material-ui/icons/AssignmentOutlined';
import InputAdornment from '@material-ui/core/InputAdornment';
import DialogActions from '@material-ui/core/DialogActions';
import Button from '@material-ui/core/Button';

interface Props {
  handleClose: () => void;
  open: boolean;
  inviteCode: string;
}
export default function InviteDialog({ open, handleClose, inviteCode }: Props) {
  return (
    <Dialog
      onClose={handleClose}
      aria-labelledby="invite-dialog-title"
      open={open}
    >
      <DialogTitle id="invite-dialog-title">Mitspieler Einladen</DialogTitle>
      <DialogContent>
        <DialogContentText>
          Über diesen persönlichen Einladungslink kannst du Freunde direkt an
          deinen Tisch hcolen:
        </DialogContentText>

        <div style={{ display: 'flex' }}>
          <Input
            fullWidth
            type="text"
            defaultValue={inviteCode}
            readOnly
            endAdornment={
              <InputAdornment position="end">
                <IconButton>
                  <AssignmentOutlinedIcon />
                </IconButton>
              </InputAdornment>
            }
          />
        </div>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose} color="primary">
          Schließen
        </Button>
      </DialogActions>
    </Dialog>
  );
}
