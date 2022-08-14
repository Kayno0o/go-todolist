import { PriorityObjectType } from '../types/PriorityType';

export const Priority: PriorityObjectType[] = [
  {
    id: 1,
    color: 'yellow',
    priority: 'low',
  },
  {
    id: 2,
    color: 'orange',
    priority: 'medium',
  },
  {
    id: 3,
    color: 'red',
    priority: 'high',
  },
];

export const getPriority = (id: number): PriorityObjectType | undefined =>
  Object.values(Priority).find((p) => p.id === id);
