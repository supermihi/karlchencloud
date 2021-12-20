import React from 'react';
import TrickView from 'play/view/TrickView';
import { Game } from 'model/match';
import { Size } from 'shared/resizeHook';

interface Props {
  game: Game;
  tableSize: Size;
}

const InGameView: React.FC<Props> = ({ game, tableSize }) => {
  const cardWidth = Math.max(tableSize.width, tableSize.height) / 5;
  return <TrickView trick={game.currentTrick} cardWidth={cardWidth} center={['50%', '50%']} />;
};
export default InGameView;
