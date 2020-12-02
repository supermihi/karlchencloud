import React from 'react';
import LoginView from 'features/auth/LoginView';
import { action } from '@storybook/addon-actions';
import { MyUserData } from 'app/auth';

/* eslint import/no-anonymous-default-export: [2, {"allowObject": true}] */
export default {
  title: 'Auth/LoginView',
  component: LoginView,
};

const login: MyUserData = { name: 'Michael', id: '123', secret: '123' };
export const WelcomeBack = () => (
  <div style={{ minWidth: '250px', minHeight: '180px' }}>
    <LoginView
      currentLogin={login}
      loading={false}
      login={action('login')}
      forgetLogin={action('forget login')}
    />
  </div>
);

export const LoginError = () => (
  <div style={{ minWidth: '250px', minHeight: '180px' }}>
    <LoginView
      currentLogin={login}
      loading={false}
      error={new Error('oh mein gott')}
      login={action('login')}
      forgetLogin={action('forget login')}
    />
  </div>
);

export const Loading = () => (
  <div style={{ minWidth: '250px', minHeight: '180px' }}>
    <LoginView
      currentLogin={login}
      loading={true}
      login={action('login')}
      forgetLogin={action('forget login')}
    />
  </div>
);
