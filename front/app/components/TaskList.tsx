import React, { useMemo } from 'react';
import { twMerge } from 'tailwind-merge';
import { Task } from '../types/TaskType';
import TaskCard from './TaskCard';

type TaskListProps = {
  tasks: Task[];
  className?: string;
};

export default function TaskList(props: TaskListProps) {
  const { tasks, className } = props;

  // sort tasks by priority
  const sortedTasks = useMemo(() => {
    return tasks
      .sort((a, b) => {
        return a.priority.id - b.priority.id;
      })
      .reverse();
  }, [tasks]);

  return (
    <div className={twMerge('flex flex-col gap-2', className)}>
      {sortedTasks.length > 0 ? (
        sortedTasks.map((task) => <TaskCard key={task.id} task={task} />)
      ) : (
        <div className="text-center text-gray-500">No tasks</div>
      )}
    </div>
  );
}
