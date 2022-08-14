import React from 'react';
import { twMerge } from 'tailwind-merge';

type FormBlockProps = {
  children: React.ReactNode;
  className?: string;
};

export default function FormBlock(props: FormBlockProps) {
  return <div className={twMerge('flex flex-col', props.className)}>{props.children}</div>;
}
