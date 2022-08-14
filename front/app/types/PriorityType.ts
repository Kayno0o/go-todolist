import { ColorType } from '../utils/Colors';

export type PriorityType = 'low' | 'medium' | 'high';
export type PriorityNumberType = 1 | 2 | 3;

export type PriorityObjectType = {
  id: PriorityNumberType;
  color: ColorType;
  priority: PriorityType;
};

export const PriorityObjects: { [name in PriorityType]: PriorityObjectType } = {
  low: { id: 1, color: 'yellow', priority: 'low' },
  medium: { id: 2, color: 'orange', priority: 'medium' },
  high: { id: 3, color: 'red', priority: 'high' },
};
