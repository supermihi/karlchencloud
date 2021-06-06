import Box from '@material-ui/core/Box';
import Tabs from '@material-ui/core/Tabs';
import React from 'react';
import { LoginData, RegisterData } from '../model';
import LoginView from './LoginView';
import RegisterView from './RegisterView';
import Tab from '@material-ui/core/Tab';
import { AppBar } from '@material-ui/core';

interface Props {
  loading: boolean;
  error?: unknown;
  login: (loginData: LoginData) => void;
  register: (registerData: RegisterData) => void;
}
interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}
function TabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`scrollable-auto-tabpanel-${index}`}
      aria-labelledby={`scrollable-auto-tab-${index}`}
      {...other}
    >
      {value === index && <Box>{children}</Box>}
    </div>
  );
}

export default function LoginOrRegisterView({
  loading,
  login,
  register,
  error,
}: Props): React.ReactElement {
  const [value, setValue] = React.useState(0);
  return (
    <>
      <AppBar position="static" color="default">
        <Tabs
          centered
          indicatorColor="primary"
          textColor="primary"
          value={value}
          onChange={(ev, x) => setValue(x as unknown as number)}
        >
          <Tab label="Login" />
          <Tab label="Register" />
        </Tabs>
      </AppBar>
      <TabPanel value={value} index={0}>
        <LoginView loading={loading} login={login} error={error} />
      </TabPanel>
      <TabPanel value={value} index={1}>
        <RegisterView loading={loading} register={register} error={error} />
      </TabPanel>
    </>
  );
}
