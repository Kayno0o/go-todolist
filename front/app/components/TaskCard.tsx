import { faCircle, faCircleExclamation } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from 'react';
import { twMerge } from 'tailwind-merge';
import { Task } from '../types/TaskType';
import { BorderColor, TextColors } from '../utils/Colors';

type TaskCardProps = {
  task: Task;
  className?: string;
};

export default function TaskCard(props: TaskCardProps) {
  const { task, className } = props;

  return (
    <div className={twMerge('rounded-full px-2 py-1 border', BorderColor[task.priority.color], className)}>
      <FontAwesomeIcon
        icon={task.priority.priority == 'high' ? faCircleExclamation : faCircle}
        className={twMerge('mr-2', TextColors[task.priority.color])}
      />
      {task.name}
    </div>
  );
}
