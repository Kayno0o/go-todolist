import { faArrowLeft } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useRouter } from 'next/router';
import React from 'react';
import { twMerge } from 'tailwind-merge';

type BackButtonProps = {
  route: string;
  className?: string;
};

export default function BackButton(props: BackButtonProps) {
  const router = useRouter();

  return (
    <FontAwesomeIcon
      icon={faArrowLeft}
      className={twMerge(
        'mr-4 cursor-pointer text-slate-700 hover:text-blue-500 transition-colors duration-300',
        props.className
      )}
      onClick={() => {
        router.push(props.route);
      }}
    />
  );
}
