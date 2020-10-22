import React from 'react';
import Button from '@material-ui/core/Button';
import { BidType, GameType } from 'api/karlchen_pb';

interface Props {
  declare: (gt: GameType) => void;
}
export default function () {
  return (
    <div>
      <Button>Gesund</Button>
    </div>
  );
}
