import React from 'react';
import TrickView from 'play/view/TrickView';
import { Game } from 'model/match';
import { Size } from 'shared/resizeHook';

interface Props {
  game: Game;
  tableSize: Size;
}

const InGameView: React.FC<Props> = ({ game, tableSize }) => {
  return (
    <TrickView trick={game.currentTrick} cardWidth={tableSize.width / 6} center={['50%', '50%']} />
  );
};
export default InGameView;
