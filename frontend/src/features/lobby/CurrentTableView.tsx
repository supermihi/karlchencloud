import React from 'react';
import { Table } from 'model/table';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardActions from '@material-ui/core/CardActions';
import Button from '@material-ui/core/Button';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import makeStyles from '@material-ui/core/styles/makeStyles';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';
import Typography from '@material-ui/core/Typography';
interface Props {
  table: Table;
}
const useStyles = makeStyles((theme) => ({
  continueButton: {
    marginLeft: 'auto',
  },
}));

export default function CurrentTableView({ table }: Props) {
  const owner = table.players.find((p) => p.id === table.owner);
  const classes = useStyles();
  return (
    <Card>
      <CardContent>
        <Typography
          variant="h4"
          component="h2"
        >{`${owner?.name}'s Tisch`}</Typography>
        <List>
          {table.players.map((player) => (
            <ListItem key={player.id}>
              <ListItemAvatar>
                <Avatar>{player.name[0].toUpperCase()}</Avatar>
              </ListItemAvatar>
              <ListItemText
                primary={player.name}
                secondary={player.online ? 'online' : 'offline'}
              ></ListItemText>
            </ListItem>
          ))}
        </List>
      </CardContent>
      <CardActions disableSpacing>
        <Button>Mitspieler einladen</Button>
        <Button
          variant="contained"
          color="primary"
          className={classes.continueButton}
          size="small"
        >
          Weiter spielen
        </Button>
      </CardActions>
    </Card>
  );
}
