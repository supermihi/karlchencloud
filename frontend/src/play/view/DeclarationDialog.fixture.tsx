import React from 'react';
import DeclarationDialog from './DeclarationDialog';
import { matchInAuction } from './mocks';

function DeclarationDialogFixture(): React.ReactElement {
  return (
    <DeclarationDialog dispatch={(action: unknown) => console.log(action)} match={matchInAuction} />
  );
}
export default DeclarationDialogFixture;
