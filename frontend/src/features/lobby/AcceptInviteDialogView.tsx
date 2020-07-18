import * as React from 'react';

import Dialog from '@material-ui/core/Dialog';
import Button from '@material-ui/core/Button';
import DialogTitle from '@material-ui/core/DialogTitle';
import DialogContent from '@material-ui/core/DialogContent';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContentText from '@material-ui/core/DialogContentText';

interface Props {
  open: boolean;
  confirm: () => void;
  reject: () => void;
}
export default function AcceptInviteDialog({ open, confirm, reject }: Props) {
  return (
    <Dialog onClose={reject} aria-labelledby="accept-invite-dialog-title" open={open}>
      <DialogTitle id="accept-invite-dialog-title">Einladung annehmen</DialogTitle>
      <DialogContent>
        <DialogContentText>Möchtest du die Tisch-Einladung annehmen?</DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button onClick={reject} color="primary">
          Nein
        </Button>
        <Button onClick={confirm} color="primary">
          Ja
        </Button>
      </DialogActions>
    </Dialog>
  );
}
