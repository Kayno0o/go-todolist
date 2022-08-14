import { faCheckCircle, faXmarkCircle } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React, { useMemo } from 'react';
import { twMerge } from 'tailwind-merge';
import { TaskGroup } from '../types/GroupList';
import { BackgroundColors, BorderColor, TextColors } from '../utils/Colors';

type GroupCardProps = {
  group: TaskGroup;
  className?: string;
  onClick?: () => void;
};

export default function GroupCard(props: GroupCardProps) {
  const { group, className } = props;

  const progress = useMemo(() => {
    const total = group.tasks.length;
    const completed = group.tasks.reduce((acc, task) => {
      if (task.type == 'quantity') {
        if (task.completed >= task.amount) return acc + 1;
        return acc + task.completed / task.amount;
      } else {
        return task.completed ? acc + 1 : acc;
      }
    }, 0);
    return Math.round((completed / total) * 100);
  }, [group.tasks]);

  const progressBg = useMemo(() => {
    if (progress > 75) return 'bg-green-500';
    if (progress > 50) return 'bg-yellow-400';
    if (progress > 25) return 'bg-orange-500';
    return 'bg-red-500';
  }, [progress]);

  const progressColor = useMemo(() => {
    if (progress > 75) return 'text-green-500';
    if (progress > 50) return 'text-yellow-400';
    if (progress > 25) return 'text-orange-500';
    return 'text-red-500';
  }, [progress]);

  return (
    <div
      className={twMerge(
        'rounded-md p-3 max-w-xs',
        'shadow-none hover:shadow-md transition-shadow duration-300',
        BackgroundColors[group.color],
        BorderColor[group.color],
        group.id ? 'cursor-pointer' : '',
        className
      )}
      onClick={props.onClick}
    >
      <FontAwesomeIcon icon={group.icon} size="2x" className={`${TextColors[group.color]}`} />

      <div className="mt-2">
        <h2 className="text-lg font-bold text-slate-800">
          {group.name} - <span className="text-gray-500 text-sm">{group.tasks.length} tasks</span>
        </h2>
      </div>

      {group.tasks.length ? (
        <div className="mt-4">
          <div className="flex items-center">
            <FontAwesomeIcon icon={progress == 100 ? faCheckCircle : faXmarkCircle} className={`${progressColor}`} />
            <span className="ml-2 text-gray-500 text-sm">{progress}%</span>
          </div>
          <div className="h-1 rounded-full mt-2 w-full bg-gray-300">
            <div
              className={twMerge('h-1 rounded-full', progressBg)}
              style={{
                width: `${progress}%`,
              }}
            />
          </div>
        </div>
      ) : null}
    </div>
  );
}
