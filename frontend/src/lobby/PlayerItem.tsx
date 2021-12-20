import * as React from 'react';
import Chip from '@material-ui/core/Chip';
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
  return (
    <Chip
      avatar={<Avatar>{player.name[0].toUpperCase()}</Avatar>}
      label={`${player.name}${me ? ' (du)' : ''}`}
    />
  );
}
