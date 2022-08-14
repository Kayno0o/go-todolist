export type ColorType = 'red' | 'orange' | 'yellow' | 'green' | 'teal' | 'blue' | 'indigo' | 'purple' | 'pink' | 'gray';

export const TextColors: { [color in ColorType]: string } = {
  red: 'text-red-500',
  orange: 'text-orange-500',
  yellow: 'text-yellow-500',
  green: 'text-green-500',
  teal: 'text-teal-500',
  blue: 'text-blue-500',
  indigo: 'text-indigo-500',
  purple: 'text-purple-500',
  pink: 'text-pink-500',
  gray: 'text-slate-500',
};

export const BackgroundColors: { [color in ColorType]: string } = {
  red: 'bg-red-50',
  orange: 'bg-orange-50',
  yellow: 'bg-yellow-50',
  green: 'bg-green-50',
  teal: 'bg-teal-50',
  blue: 'bg-blue-50',
  indigo: 'bg-indigo-50',
  purple: 'bg-purple-50',
  pink: 'bg-pink-50',
  gray: 'bg-gray-50',
};

export const BrightBackgroundColors: { [color in ColorType]: string } = {
  red: 'bg-red-400',
  orange: 'bg-orange-400',
  yellow: 'bg-yellow-400',
  green: 'bg-green-400',
  teal: 'bg-teal-400',
  blue: 'bg-blue-400',
  indigo: 'bg-indigo-400',
  purple: 'bg-purple-400',
  pink: 'bg-pink-400',
  gray: 'bg-gray-400',
};

export const BorderColor: { [color in ColorType]: string } = {
  red: 'border border-red-300',
  orange: 'border border-orange-300',
  yellow: 'border border-yellow-300',
  green: 'border border-green-300',
  teal: 'border border-teal-300',
  blue: 'border border-blue-300',
  indigo: 'border border-indigo-300',
  purple: 'border border-purple-300',
  pink: 'border border-pink-300',
  gray: 'border border-gray-300',
};
