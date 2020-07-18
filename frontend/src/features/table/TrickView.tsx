import React from 'react';
import { Trick } from 'model/match';

interface Props {
  trick: Trick;
}

export default function TrickView({ trick }: Props) {
  return 'aktueller Stich';
}
