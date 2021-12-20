import React from 'react';
import AppBar from './AppBar';
import Container from '@material-ui/core/Container';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';

export default (
  <div style={{ display: 'flex', flexDirection: 'column' }}>
    <AppBar location="Spieltische" />
    <Container style={{ marginTop: 6, flexGrow: 1 }}>
      <Paper color="secondary">
        <Typography variant="h1">bla</Typography>
      </Paper>
    </Container>
  </div>
);
