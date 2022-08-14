import { IconProp } from '@fortawesome/fontawesome-svg-core';
import { ColorType } from '../utils/Colors';
import { IconType } from '../utils/Icons';
import { Task } from './TaskType';

export type TaskGroup = {
  id?: number;
  name: string;
  description: string;
  tasks: Task[];
  icon: IconProp;
  color: ColorType;
};

export type JSONGroup = {
  id: 1;
  name: string;
  description: string;
  icon: IconType;
  color: ColorType;
};
