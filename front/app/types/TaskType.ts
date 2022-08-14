import { PriorityNumberType, PriorityObjectType } from './PriorityType';

export type TaskType = 'tick' | 'quantity';

export type Task = {
  id?: number;
  name: string;
  priority: PriorityObjectType;
  task_group: number;
  type: TaskType;
} & (
  | {
      type: 'tick';
      completed: boolean;
    }
  | {
      type: 'quantity';
      completed: number;
      amount: number;
    }
);

export type JSONTask = {
  id: number;
  name: string;
  completed: boolean | number;
  priority: PriorityNumberType;
  type: TaskType;
  task_group: number;
  amount: number;
};
