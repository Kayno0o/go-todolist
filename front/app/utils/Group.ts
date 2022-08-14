import { JSONGroup, TaskGroup } from '../types/GroupList';
import { JSONTask } from '../types/TaskType';
import { TaskIcons } from './Icons';
import { JsonToTask } from './Task';

export async function JsonToGroup(group: JSONGroup): Promise<TaskGroup> {
  const tasks: JSONTask[] = await fetch(`/api/task/all/task_group/${group.id}`).then((res) => res.json());

  return {
    id: group.id,
    name: group.name,
    description: group.description,
    color: group.color,
    icon: TaskIcons[group.icon],
    tasks: tasks.map(JsonToTask),
  };
}
