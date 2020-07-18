import React from 'react';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import Typography from '@material-ui/core/Typography';
import { User } from 'model/core';

interface Props {
  user: User;
}

export default function PlayerView({ user }: Props) {
  return (
    <Card variant="outlined">
      <CardHeader title={user.name}></CardHeader>
      <CardContent>
        <Typography variant="body2">Re Contra bla überlegt</Typography>
      </CardContent>
    </Card>
  );
}
