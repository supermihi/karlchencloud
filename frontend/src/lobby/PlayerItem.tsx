import * as React from 'react';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';
import makeStyles from '@material-ui/core/styles/makeStyles';
import { User } from 'model/core';

interface Props {
  player: User;
  me: boolean;
}

const useStyles = makeStyles((theme) => ({
  online: {
    color: theme.palette.success.main,
  } as const,
}));

export default function PlayerItem({ player, me }: Props): React.ReactElement {
  const classes = useStyles();
  return (
    <ListItem>
      <ListItemAvatar>
        <Avatar>{player.name[0].toUpperCase()}</Avatar>
      </ListItemAvatar>
      <ListItemText
        primary={player.name}
        secondary={me ? 'du' : player.online && <em className={classes.online}>online</em>}
      />
    </ListItem>
  );
}
