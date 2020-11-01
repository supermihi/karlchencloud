import React from 'react';
import Dialog from '@material-ui/core/Dialog';
import DialogTitle from '@material-ui/core/DialogTitle';
import { GameType } from 'api/karlchen_pb';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import makeStyles from '@material-ui/core/styles/makeStyles';

const useStyles = makeStyles((theme) => ({
  root: {
    backgroundColor: theme.palette.background.paper,
  },
}));
interface Props {
  declare: (gt: GameType) => void;
  open: boolean;
}
export default function ({ declare, open }: Props) {
  const classes = useStyles();
  return (
    <Dialog open={open}>
      <DialogTitle>Deine Ansage</DialogTitle>
      <List className={classes.root}>
        <ListItem button>
          <ListItemText primary="Gesund" onClick={() => declare(GameType.NORMAL_GAME)} />
        </ListItem>
        <ListItem button>
          <ListItemText primary="Hochzeit" onClick={() => declare(GameType.MARRIAGE)} />
        </ListItem>
      </List>
    </Dialog>
  );
}
