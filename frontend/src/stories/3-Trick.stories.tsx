import React from 'react';
import TrickView from 'features/table/TrickView';
import {Diamond10, Diamond9, DiamondA, DiamondQ} from "../model/cards";

export default {
  title: 'Match/Trick',
  component: TrickView,
};
export const Card = () => (
  <div style={{height: "250px"}}>
  <TrickView
    cardHeight="8cm"
    trick={{forehand: 'pl0', cards: {pl0: Diamond9, pl1: Diamond10, pl2: DiamondA, me: DiamondQ} }}
    players={['me', 'pl0', 'pl1', 'pl2']}
  />
  </div>
);
