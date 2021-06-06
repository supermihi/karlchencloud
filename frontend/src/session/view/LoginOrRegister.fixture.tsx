import React from 'react';
import { useValue } from 'react-cosmos/fixture';
import { mockGrpcError } from 'shared/mock';
import LoginOrRegisterView from './LoginOrRegisterView';

const Fixture: React.FC = () => {
  const [error] = useValue('error', { defaultValue: false });
  const [loading] = useValue('loading', { defaultValue: false });
  return (
    <LoginOrRegisterView
      loading={loading}
      error={error && mockGrpcError('mock')}
      login={() => console.log('login')}
      register={() => console.log('register')}
    />
  );
};
export default Fixture;
