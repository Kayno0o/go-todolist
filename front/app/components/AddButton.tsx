import { faPlusCircle } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from 'react';
import { twMerge } from 'tailwind-merge';

type AddButtonProps = {
  className?: string;
  onClick: () => void;
};

export default function AddButton(props: AddButtonProps) {
  return (
    <div className={twMerge('fixed bottom-0 right-0 m-4 cursor-pointer', props.className)} onClick={props.onClick}>
      <FontAwesomeIcon
        icon={faPlusCircle}
        className="text-blue-500 rotate-0 hover:rotate-90 transition-transform duration-300"
        size="3x"
      />
    </div>
  );
}
