import React from 'react';
import makeStyles from '@material-ui/core/styles/makeStyles';
import Backdrop from '@material-ui/core/Backdrop';
import CircularProgress from '@material-ui/core/CircularProgress';

interface Props {
  open: boolean;
}
function SpinBackdrop({ open }: Props): React.ReactElement {
  const classes = useStyles();
  return (
    <Backdrop open={open} className={classes.backdrop}>
      <CircularProgress />
    </Backdrop>
  );
}
export default SpinBackdrop;

const useStyles = makeStyles((theme) => ({
  backdrop: {
    zIndex: theme.zIndex.drawer + 1,
    color: '#fff',
  },
}));
