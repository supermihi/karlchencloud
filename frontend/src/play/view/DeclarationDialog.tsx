import React from 'react';
import Dialog from '@material-ui/core/Dialog';
import DialogTitle from '@material-ui/core/DialogTitle';
import { GameType, MatchPhase } from 'api/karlchen_pb';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import makeStyles from '@material-ui/core/styles/makeStyles';
import { AppDispatch } from '../../state';
import { declare } from '../thunks';
import { Match } from '../../model/match';
import { Pos } from '../../model/players';

const useStyles = makeStyles((theme) => ({
  root: {
    backgroundColor: theme.palette.background.paper,
  },
}));

interface Props {
  dispatch: AppDispatch;
  match: Match;
}

export default function DeclarationDialog({ dispatch, match }: Props): React.ReactElement {
  const open = match.phase === MatchPhase.AUCTION && match.turn === Pos.bottom;
  const classes = useStyles();
  return (
    <Dialog open={open}>
      <DialogTitle>Deine Ansage</DialogTitle>
      <List className={classes.root}>
        <ListItem button>
          <ListItemText primary="Gesund" onClick={() => dispatch(declare(GameType.NORMAL_GAME))} />
        </ListItem>
        <ListItem button>
          <ListItemText primary="Hochzeit" onClick={() => dispatch(declare(GameType.MARRIAGE))} />
        </ListItem>
      </List>
    </Dialog>
  );
}
