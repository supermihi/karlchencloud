import * as React from 'react';
import Toolbar from '@material-ui/core/Toolbar';
import makeStyles from '@material-ui/core/styles/makeStyles';
import Typography from '@material-ui/core/Typography';
import AppBar from '@material-ui/core/AppBar';
import IconButton from '@material-ui/core/IconButton';
import AccountCircle from '@material-ui/icons/AccountCircle';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import GrowDiv from '../components/GrowDiv';

const useStyles = makeStyles(() => ({
  appBar: {
    position: 'relative',
    display: 'flex',
  } as const,
}));
interface Props {
  loggedIn: boolean;
  logout: () => void;
}
export default function MyAppBar({ loggedIn, logout }: Props) {
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const classes = useStyles();
  return (
    <AppBar position="absolute" className={classes.appBar}>
      <Toolbar>
        <Typography variant="h6" color="inherit" component="h1" noWrap>
          Karlchencloud
        </Typography>
        {loggedIn && (
          <>
            <GrowDiv />

            <IconButton
              aria-controls="account-menu"
              aria-haspopup="true"
              edge="end"
              color="inherit"
              onClick={(e) => setAnchorEl(e.currentTarget)}
            >
              <AccountCircle />
            </IconButton>
            <Menu
              id="account-menu"
              anchorEl={anchorEl}
              keepMounted
              open={Boolean(anchorEl)}
              onClose={() => {
                setAnchorEl(null);
              }}
            >
              <MenuItem
                onClick={() => {
                  logout();
                  setAnchorEl(null);
                }}
              >
                Logout
              </MenuItem>
            </Menu>
          </>
        )}
      </Toolbar>
    </AppBar>
  );
}
