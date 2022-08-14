import React from 'react';
import { twMerge } from 'tailwind-merge';
import FormBlock from './FormBlock';

type InputProps = {
  className?: string;
  name: string;
  label: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => void;
  type?: string;
};

export default function Input(props: InputProps) {
  return (
    <FormBlock>
      <label htmlFor={`input-${props.name}`} className={'text-gray-700'}>
        {props.label}
      </label>

      {props.type === 'textarea' ? (
        <textarea
          className={twMerge('border border-gray-300 rounded-lg p-2', props.className)}
          name={props.name}
          id={`input-${props.name}`}
          value={props.value}
          onChange={props.onChange}
        />
      ) : (
        <input
          className={twMerge('border border-gray-300 rounded-lg p-2', props.className)}
          type={props.type || 'text'}
          name={props.name}
          id={`input-${props.name}`}
          value={props.value}
          onChange={props.onChange}
        />
      )}
    </FormBlock>
  );
}
