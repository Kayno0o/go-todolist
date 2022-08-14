import { faBars } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from 'react';

export function Header() {
  return (
    <header className="p-4 fixed bg-white w-full">
      <FontAwesomeIcon icon={faBars} className="text-slate-800" size="2x" />
    </header>
  );
}
