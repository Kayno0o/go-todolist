import React from 'react';
import { twMerge } from 'tailwind-merge';
import FormBlock from './FormBlock';

type SelectProps = {
  className?: string;
  name: string;
  label: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLSelectElement>) => void;
  options: {
    value: string;
    label: string;
  }[];
};

export default function Select(props: SelectProps) {
  return (
    <FormBlock>
      <label htmlFor={`input-${props.name}`} className={'text-gray-700'}>
        {props.label}
      </label>
      <select
        className={twMerge('border border-gray-300 bg-white rounded-lg p-2', props.className)}
        name={props.name}
        id={`input-${props.name}`}
        value={props.value}
        onChange={props.onChange}
      >
        {props.options.map((option) => (
          <option key={option.value} value={option.value}>
            {option.label}
          </option>
        ))}
      </select>
    </FormBlock>
  );
}
