import React from 'react';
import Card, { CardProps } from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import { User } from 'model/core';

interface Props extends CardProps {
  user: User;
}

export default function PlayerView({ user, ...props }: Props) {
  return (
    <Card variant="outlined" {...props}>
      <CardHeader titleTypographyProps={{ variant: 'h6' }} title={user.name}></CardHeader>
      <CardContent></CardContent>
    </Card>
  );
}
